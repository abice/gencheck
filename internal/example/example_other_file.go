package example

import (
	"errors"

	"github.com/abice/gencheck"
)

// OtherFile is another example struct for testing
type OtherFile struct {
	EqCSFieldString  string `valid:"required"`
	NeCSFieldString  string
	GtCSFieldString  string
	GteCSFieldString string
	LtCSFieldString  string
	LteCSFieldString string
}

func (s OtherFile) Validate() error {

	vErrors := make(gencheck.ValidationErrors, 0, 1)

	// BEGIN EqCSFieldString Validations
	// required
	if s.EqCSFieldString == "" {
		vErrors = append(vErrors, gencheck.NewFieldError("OtherFile", "EqCSFieldString", "required", errors.New("is required")))
	}
	// END EqCSFieldString Validations

	if len(vErrors) > 0 {
		return vErrors
	}

	return nil
}
