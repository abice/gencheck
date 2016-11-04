package gencheck

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	errMsgFormat = `validation: field validation failed for '%s.%s' on rule '%s'`
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
	tag        string
	field      string
	structName string
	msg        string
}

func (fe fieldError) Error() string {
	msg := fe.msg
	if msg == "" {
		msg = fe.tag
	}
	return fmt.Sprintf(errMsgFormat, fe.structName, fe.field, msg)
}

func (fe fieldError) Tag() string {
	return fe.tag
}

func (fe fieldError) Field() string {
	return fe.field
}

func (fe fieldError) Struct() string {
	return fe.structName
}

func (fe fieldError) Message() string {
	return fe.msg
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
		structName: st,
		field:      field,
		tag:        tag,
		msg:        msg,
	}
}
