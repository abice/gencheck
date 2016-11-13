package example

import (
	"errors"
	"fmt"
	"time"

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

	vErrors := make(gencheck.ValidationErrors, 0, 31)

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
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "MinString", "min", errors.New("length failed check for min=1")))
	}

	// END MinString Validations

	// BEGIN MinNumber Validations
	// min

	if s.MinNumber < 1113.00 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "MinNumber", "min", errors.New("failed check for min=1113.00")))
	}

	// END MinNumber Validations

	// BEGIN MinMultiple Validations
	// min

	if len(s.MinMultiple) < 7 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "MinMultiple", "min", errors.New("length failed check for min=7")))
	}

	// END MinMultiple Validations

	// BEGIN MaxString Validations
	// max

	if len(s.MaxString) > 3 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "MaxString", "max", errors.New("length failed check for max=3")))
	}

	// END MaxString Validations

	// BEGIN MaxNumber Validations
	// max

	if s.MaxNumber > 1113.00 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "MaxNumber", "max", errors.New("failed check for max=1113.00")))
	}

	// END MaxNumber Validations

	// BEGIN MaxMultiple Validations
	// max

	if len(s.MaxMultiple) > 7 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "MaxMultiple", "max", errors.New("length failed check for max=7")))
	}

	// END MaxMultiple Validations

	// BEGIN LtString Validations
	// lt

	if len(s.LtString) >= 3 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "LtString", "lt", errors.New("length failed check for lt=3")))
	}

	// END LtString Validations

	// BEGIN LtNumber Validations
	// lt

	if s.LtNumber >= 5.56 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "LtNumber", "lt", errors.New("failed check for lt=5.56")))
	}

	// END LtNumber Validations

	// BEGIN LtMultiple Validations
	// lt

	if len(s.LtMultiple) >= 2 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "LtMultiple", "lt", errors.New("length failed check for lt=2")))
	}

	// END LtMultiple Validations

	// BEGIN LtTime Validations
	// lt

	tLtTime := time.Now().UTC()
	if !s.LtTime.Before(tLtTime) {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "LtTime", "lt", errors.New("is after now")))
	}

	// END LtTime Validations

	// BEGIN LteString Validations
	// lte

	if len(s.LteString) > 3 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "LteString", "lte", errors.New("length failed check for lte=3")))
	}

	// END LteString Validations

	// BEGIN LteNumber Validations
	// lte

	if s.LteNumber > 5.56 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "LteNumber", "lte", errors.New("failed check for lte=5.56")))
	}

	// END LteNumber Validations

	// BEGIN LteMultiple Validations
	// lte

	if len(s.LteMultiple) > 2 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "LteMultiple", "lte", errors.New("length failed check for lte=2")))
	}

	// END LteMultiple Validations

	// BEGIN LteTime Validations
	// lte

	tLteTime := time.Now().UTC()
	if s.LteTime.After(tLteTime) {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "LteTime", "lte", errors.New("is after now")))
	}

	// END LteTime Validations

	// BEGIN GtString Validations
	// gt

	if len(s.GtString) <= 3 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "GtString", "gt", errors.New("length failed check for gt=3")))
	}

	// END GtString Validations

	// BEGIN GtNumber Validations
	// gt

	if s.GtNumber <= 5.56 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "GtNumber", "gt", errors.New("failed check for gt=5.56")))
	}

	// END GtNumber Validations

	// BEGIN GtMultiple Validations
	// gt

	if len(s.GtMultiple) <= 2 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "GtMultiple", "gt", errors.New("length failed check for gt=2")))
	}

	// END GtMultiple Validations

	// BEGIN GtTime Validations
	// gt

	tGtTime := time.Now().UTC()
	if !s.GtTime.After(tGtTime) {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "GtTime", "gt", errors.New("is before now")))
	}

	// END GtTime Validations

	// BEGIN GteString Validations
	// gte

	if len(s.GteString) < 3 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "GteString", "gte", errors.New("length failed check for gte=3")))
	}

	// END GteString Validations

	// BEGIN GteNumber Validations
	// gte

	if s.GteNumber < 5.56 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "GteNumber", "gte", errors.New("failed check for gte=5.56")))
	}

	// END GteNumber Validations

	// BEGIN GteMultiple Validations
	// gte

	if len(s.GteMultiple) < 2 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "GteMultiple", "gte", errors.New("length failed check for gte=2")))
	}

	// END GteMultiple Validations

	// BEGIN GteTime Validations
	// gte

	tGteTime := time.Now().UTC()
	if s.GteTime.Before(tGteTime) {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "GteTime", "gte", errors.New("is before now")))
	}

	// END GteTime Validations

	// BEGIN GteTimeVal Validations
	// gte

	tGteTimeVal := time.Now().UTC().Add(1 * time.Second)
	if s.GteTimeVal.Before(tGteTimeVal) {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "GteTimeVal", "gte", fmt.Errorf("is before %s", tGteTimeVal)))
	}

	// END GteTimeVal Validations

	// BEGIN GteTimePtr Validations
	// gte

	if s.GteTimePtr != nil {
		tGteTimePtr := time.Now().UTC()
		if (*s.GteTimePtr).Before(tGteTimePtr) {
			vErrors = append(vErrors, gencheck.NewFieldError("Test", "GteTimePtr", "gte", errors.New("is before now")))
		}
	}

	// END GteTimePtr Validations

	// BEGIN UUID Validations
	// uuid

	if err := gencheck.IsUUID(&s.UUID); err != nil {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "UUID", "uuid", err))
	}

	// END UUID Validations

	// BEGIN MinIntPtr Validations
	// required

	if s.MinIntPtr == nil {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "MinIntPtr", "required", errors.New("is required")))
	}

	// min

	if s.MinIntPtr != nil && *s.MinIntPtr < 1234 {
		vErrors = append(vErrors, gencheck.NewFieldError("Test", "MinIntPtr", "min", errors.New("failed check for min=1234")))
	}

	// END MinIntPtr Validations

	if len(vErrors) > 0 {
		return vErrors
	}

	return nil
}
