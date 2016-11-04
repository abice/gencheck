package example

import (
	"errors"

	"github.com/abice/gencheck"
)

func (s Example) Validate() error {
	var vErrors gencheck.ValidationErrors

	// BEGIN MapOfInterfaces Validations
	// notnil

	if s.MapOfInterfaces == nil {
		vErrors = append(vErrors, gencheck.NewFieldError("Example", "MapOfInterfaces", "notnil", errors.New("is Nil")))
	}

	// END MapOfInterfaces Validations

	if len(vErrors) > 0 {
		return vErrors
	}
	return nil
}

func (s Test) Validate() error {
	var vErrors gencheck.ValidationErrors

	// BEGIN LenMultiple Validations
	// len

	if !(len(s.LenMultiple) == 7) {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "LenMultiple", "len", errors.New("length mismatch")))
	}

	// END LenMultiple Validations

	// BEGIN LenNumber Validations
	// len

	if !(s.LenNumber == 1113.00) {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "LenNumber", "len", errors.New("length mismatch")))
	}

	// END LenNumber Validations

	// BEGIN LenString Validations
	// len

	if !(len(s.LenString) == 1) {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "LenString", "len", errors.New("length mismatch")))
	}

	// END LenString Validations

	// BEGIN RequiredMultiple Validations
	// required

	if s.RequiredMultiple == nil {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "RequiredMultiple", "required", errors.New("is required")))
	}

	// END RequiredMultiple Validations

	// BEGIN RequiredString Validations
	// required

	if s.RequiredString == "" {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "RequiredString", "required", errors.New("is required")))
	}

	// END RequiredString Validations

	// BEGIN UUID Validations
	// uuid

	if err := gencheck.IsUUID(&s.UUID); err != nil {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "UUID", "uuid", err))
	}

	// END UUID Validations

	if len(vErrors) > 0 {
		return vErrors
	}
	return nil
}
