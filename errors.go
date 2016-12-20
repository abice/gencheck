package gencheck

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	errMsgFormat = `validation: field validation failed for '%s.%s': %s`
	errTagFormat = `validation: field validation failed for '%s.%s': tag='%s'`
)

// Error returns a string of the validation errors
func (ve ValidationErrors) Error() string {
	buff := bytes.NewBufferString("")

	var fe *fieldError

	for i := 0; i < len(ve); i++ {

		fe = ve[i].(*fieldError)
		buff.WriteString(fe.Error())
		buff.WriteString("\n")
	}

	return strings.TrimSpace(buff.String())
}

// ValidationErrors is an array of Field Errors
type ValidationErrors []FieldError

// FieldError provide error information for a specific field validation.
type FieldError interface {
	Tag() string
	Field() string
	Struct() string
	Message() string
	error
}

// fieldError provide error information for a specific field validation.
type fieldError struct {
	Ftag        string `json:"tag"`
	Ffield      string `json:"field"`
	FstructName string `json:"structName"`
	Fmsg        string `json:"msg,omitempty"`
}

func (fe fieldError) Error() string {
	if fe.Fmsg == "" {
		return fmt.Sprintf(errTagFormat, fe.FstructName, fe.Ffield, fe.Ftag)
	}
	return fmt.Sprintf(errMsgFormat, fe.FstructName, fe.Ffield, fe.Fmsg)
}

func (fe fieldError) Tag() string {
	return fe.Ftag
}

func (fe fieldError) Field() string {
	return fe.Ffield
}

func (fe fieldError) Struct() string {
	return fe.FstructName
}

func (fe fieldError) Message() string {
	return fe.Fmsg
}

// NewFieldError returns a newly created immutable FieldError
// Tag is normally the rule that was invalid, but can be used for whatever
// message you want.
func NewFieldError(st, field, tag string, err error) FieldError {
	var msg string
	if err != nil {
		msg = err.Error()
	}
	return &fieldError{
		FstructName: st,
		Ffield:      field,
		Ftag:        tag,
		Fmsg:        msg,
	}
}
