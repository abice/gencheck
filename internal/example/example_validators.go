package example

import (
	"errors"

	"github.com/abice/gencheck"
)

// Validate is an automatically generated validation method provided by
// gencheck.
// See https://github.com/abice/gencheck for more details.
func (s Example) Validate() error {

	vErrors := make(gencheck.ValidationErrors, 0, 1)

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

// Validate is an automatically generated validation method provided by
// gencheck.
// See https://github.com/abice/gencheck for more details.
func (s Test) Validate() error {

	vErrors := make(gencheck.ValidationErrors, 0, 12)

	// BEGIN RequiredString Validations
	// required

	if s.RequiredString == "" {
		return append(vErrors, gencheck.NewFieldError("Test", "RequiredString", "required", errors.New("is required")))
	}

	// END RequiredString Validations

	// BEGIN RequiredMultiple Validations
	// required

	if s.RequiredMultiple == nil {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "RequiredMultiple", "required", errors.New("is required")))
	}

	// END RequiredMultiple Validations

	// BEGIN LenString Validations
	// len

	if !(len(s.LenString) == 1) {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "LenString", "len", errors.New("length mismatch")))
	}

	// END LenString Validations

	// BEGIN LenNumber Validations
	// len

	if !(s.LenNumber == 1113.00) {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "LenNumber", "len", errors.New("length mismatch")))
	}

	// END LenNumber Validations

	// BEGIN LenMultiple Validations
	// len

	if !(len(s.LenMultiple) == 7) {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "LenMultiple", "len", errors.New("length mismatch")))
	}

	// END LenMultiple Validations

	// BEGIN MinString Validations
	// min

	if len(s.MinString) < 1 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "MinString", "min", errors.New("length was less than 1")))
	}

	// END MinString Validations

	// BEGIN MinNumber Validations
	// min

	if s.MinNumber < 1113.00 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "MinNumber", "min", errors.New("was less than 1113.00")))
	}

	// END MinNumber Validations

	// BEGIN MinMultiple Validations
	// min

	if len(s.MinMultiple) < 7 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "MinMultiple", "min", errors.New("length was less than 7")))
	}

	// END MinMultiple Validations

	// BEGIN MaxString Validations
	// max

	if len(s.MaxString) > 3 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "MaxString", "max", errors.New("length was more than 3")))
	}

	// END MaxString Validations

	// BEGIN MaxNumber Validations
	// max

	if s.MaxNumber > 1113.00 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "MaxNumber", "max", errors.New("was more than 1113.00")))
	}

	// END MaxNumber Validations

	// BEGIN MaxMultiple Validations
	// max

	if len(s.MaxMultiple) > 7 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "MaxMultiple", "max", errors.New("length was more than 7")))
	}

	// END MaxMultiple Validations

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
