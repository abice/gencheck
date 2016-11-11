package example

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/abice/gencheck"
	"github.com/stretchr/testify/suite"
)

// ExampleTestSuite: Run this test suite AFTER running gokay against example.go
type ExampleTestSuite struct {
	suite.Suite
}

// TestExampleTestSuite
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

// TestValidateTestStruct_NoValues
func (s *ExampleTestSuite) TestValidateTestStruct_FailFast() {
	expected := gencheck.ValidationErrors{
		gencheck.NewFieldError("Test", "RequiredString", "required", fmt.Errorf("is required")),
	}

	underTest := Test{
		MaxString:   "1234",
		MaxNumber:   1113.00001,
		MaxMultiple: []string{"", "", "", "", "", "", "", "", ""},
	}

	err := underTest.Validate()
	s.Require().IsType(gencheck.ValidationErrors{}, err, "Error returned was not ValidationErrors type")

	ve := err.(gencheck.ValidationErrors)

	s.Require().Equal(len(expected), len(ve), "Validation Errors were of different length")

	for i, fe := range ve {
		s.Equal(expected[i], fe)
	}
}

// TestValidateTestStruct_NoValues
func (s *ExampleTestSuite) TestValidateTestStruct_NoValues() {
	expected := gencheck.ValidationErrors{
		gencheck.NewFieldError("Test", "RequiredMultiple", "required", fmt.Errorf("is required")),
		gencheck.NewFieldError("Test", "LenString", "len", fmt.Errorf("length mismatch")),
		gencheck.NewFieldError("Test", "LenNumber", "len", fmt.Errorf("length mismatch")),
		gencheck.NewFieldError("Test", "LenMultiple", "len", fmt.Errorf("length mismatch")),
		gencheck.NewFieldError("Test", "MinString", "min", fmt.Errorf("length failed check for min=1")),
		gencheck.NewFieldError("Test", "MinNumber", "min", fmt.Errorf("failed check for min=1113.00")),
		gencheck.NewFieldError("Test", "MinMultiple", "min", fmt.Errorf("length failed check for min=7")),
		gencheck.NewFieldError("Test", "MaxString", "max", fmt.Errorf("length failed check for max=3")),
		gencheck.NewFieldError("Test", "MaxNumber", "max", fmt.Errorf("failed check for max=1113.00")),
		gencheck.NewFieldError("Test", "MaxMultiple", "max", fmt.Errorf("length failed check for max=7")),
		gencheck.NewFieldError("Test", "LtString", "lt", fmt.Errorf("length failed check for lt=3")),
		gencheck.NewFieldError("Test", "LtNumber", "lt", fmt.Errorf("failed check for lt=5.56")),
		gencheck.NewFieldError("Test", "LtMultiple", "lt", fmt.Errorf("length failed check for lt=2")),
		gencheck.NewFieldError("Test", "LtTime", "lt", fmt.Errorf("is after now")),
		gencheck.NewFieldError("Test", "LteString", "lte", fmt.Errorf("length failed check for lte=3")),
		gencheck.NewFieldError("Test", "LteNumber", "lte", fmt.Errorf("failed check for lte=5.56")),
		gencheck.NewFieldError("Test", "LteMultiple", "lte", fmt.Errorf("length failed check for lte=2")),
		gencheck.NewFieldError("Test", "GtString", "gt", fmt.Errorf("length failed check for gt=3")),
		gencheck.NewFieldError("Test", "GtNumber", "gt", fmt.Errorf("failed check for gt=5.56")),
		gencheck.NewFieldError("Test", "GtMultiple", "gt", fmt.Errorf("length failed check for gt=2")),
		gencheck.NewFieldError("Test", "GtTime", "gt", fmt.Errorf("is before now")),
		gencheck.NewFieldError("Test", "GteString", "gte", fmt.Errorf("length failed check for gte=3")),
		gencheck.NewFieldError("Test", "GteNumber", "gte", fmt.Errorf("failed check for gte=5.56")),
		gencheck.NewFieldError("Test", "GteMultiple", "gte", fmt.Errorf("length failed check for gte=2")),
		gencheck.NewFieldError("Test", "GteTime", "gte", fmt.Errorf("is before now")),
		gencheck.NewFieldError("Test", "GteTimeVal", "gte", fmt.Errorf("is before %s", time.Now().UTC().Add(1*time.Second).Truncate(time.Second).Format("2006-01-02 15:04:05"))),
		gencheck.NewFieldError("Test", "GteTimePtr", "gte", fmt.Errorf("is before now")),
		gencheck.NewFieldError("Test", "UUID", "uuid", fmt.Errorf("'' is not a UUID")),
		gencheck.NewFieldError("Test", "MinIntPtr", "required", fmt.Errorf("is required")),
	}
	testTime := time.Now().UTC()
	underTest := Test{
		RequiredString: "Here I am", // Put in required string to prevent fast failure
		MaxString:      "1234",
		MaxNumber:      1113.00001,
		MaxMultiple:    []string{"", "", "", "", "", "", "", "", ""},
		LtString:       "1234",
		LtNumber:       5.56000001,
		LtMultiple:     []string{"", "", ""},
		LtTime:         time.Now().UTC().Add(1 * time.Millisecond),
		LteString:      "1234",
		LteNumber:      5.56000001,
		LteMultiple:    []string{"", "", ""},
		GteTimePtr:     &testTime,
	}

	err := underTest.Validate()
	s.Error(err, "Error should not be nil")
	s.Require().IsType(gencheck.ValidationErrors{}, err, "Error returned was not ValidationErrors type")

	ve := err.(gencheck.ValidationErrors)

	s.Require().Equal(len(expected), len(ve), "Validation Errors were of different length")

	for i, fe := range ve {
		s.Contains(fe.Error(), strings.TrimSuffix(expected[i].Error(), "'"), "Error string mismatch")
	}
}

// TestValidateTestStruct_Values
func (s *ExampleTestSuite) TestValidateTestStruct_Values() {
	i := int64(2000)
	underTest := Test{
		LenMultiple:      []string{"", "", "", "", "", "", ""},
		LenNumber:        1113,
		LenString:        "a",
		RequiredMultiple: []string{},
		RequiredString:   "b",
		MinString:        "1234567",
		MinNumber:        1113.000001,
		MinMultiple:      []string{"", "", "", "", "", "", "", ""},
		UUID:             "7112EE37-3219-4A26-BA01-1D230BC9257B",
		MinIntPtr:        &i,
		GteString:        "1234",
		GteNumber:        5.5600001,
		GteMultiple:      []string{"", ""},
		LteTime:          time.Now().Add(-5 * time.Second),
		GtTime:           time.Now().Add(1 * time.Millisecond),
		GteTime:          time.Now().Add(5 * time.Second),
		GteTimeVal:       time.Now().Add(5 * time.Second),
		GtNumber:         5.5600001,
		GtMultiple:       []string{"", "", ""},
		GtString:         "1234",
	}

	err := underTest.Validate()
	s.NoError(err, "Valid Struct should not have had an error")

}

// TestValidateTestStruct_Values
func (s *ExampleTestSuite) TestValidateTestStruct_MinPtrFailure() {
	i := int64(1233)
	underTest := Test{
		LenMultiple:      []string{"", "", "", "", "", "", ""},
		LenNumber:        1113,
		LenString:        "a",
		RequiredMultiple: []string{},
		RequiredString:   "b",
		MinString:        "1234567",
		MinNumber:        1113.000001,
		MinMultiple:      []string{"", "", "", "", "", "", "", ""},
		UUID:             "7112EE37-3219-4A26-BA01-1D230BC9257B",
		GteString:        "1234",
		GteNumber:        5.5600001,
		GteMultiple:      []string{"", ""},
		LteTime:          time.Now().Add(-5 * time.Second),
		GtTime:           time.Now().Add(1 * time.Millisecond),
		GteTime:          time.Now().Add(5 * time.Second),
		GteTimeVal:       time.Now().Add(5 * time.Second),
		GtNumber:         5.5600001,
		GtMultiple:       []string{"", "", ""},
		GtString:         "1234",
		MinIntPtr:        &i,
	}

	err := underTest.Validate()
	s.Error(err, "Valid Struct should not have had an error")
	s.Require().IsType(gencheck.ValidationErrors{}, err, "Error returned was not ValidationErrors type")

	ve := err.(gencheck.ValidationErrors)
	s.Require().Len(ve, 1, "Should only have 1 validation error")

	s.Require().EqualValues(ve[0], gencheck.NewFieldError("Test", "MinIntPtr", "min", fmt.Errorf("failed check for min=1234")), "Error should be min error")
}

// TestValidateExample_NilMap
func (s *ExampleTestSuite) TestValidateExample_NilMap() {
	expected := gencheck.ValidationErrors{
		gencheck.NewFieldError("Example", "MapOfInterfaces", "notnil", fmt.Errorf("is Nil")),
	}

	underTest := Example{}

	err := underTest.Validate()
	s.EqualValues(expected, err)
}

// TestValidateNotNil_Map
func (s *ExampleTestSuite) TestValidateExample_Happy() {

	underTest := Example{
		MapOfInterfaces: make(map[string]interface{}),
	}

	err := underTest.Validate()
	s.Nil(err, "Valid Struct should not have had an error")
}

// TestValidateTestStruct_Values
func (s *ExampleTestSuite) TestValidateTestStruct_LteTime() {
	i := int64(1234)
	underTest := Test{
		LenMultiple:      []string{"", "", "", "", "", "", ""},
		LenNumber:        1113,
		LenString:        "a",
		RequiredMultiple: []string{},
		RequiredString:   "b",
		MinString:        "1234567",
		MinNumber:        1113.000001,
		MinMultiple:      []string{"", "", "", "", "", "", "", ""},
		UUID:             "7112EE37-3219-4A26-BA01-1D230BC9257B",
		GteString:        "1234",
		GteNumber:        5.5600001,
		GteMultiple:      []string{"", ""},
		LteTime:          time.Now().Add(1 * time.Second),
		GtTime:           time.Now().Add(1 * time.Millisecond),
		GteTime:          time.Now().Add(5 * time.Second),
		GteTimeVal:       time.Now().Add(5 * time.Second),
		GtNumber:         5.5600001,
		GtMultiple:       []string{"", "", ""},
		GtString:         "1234",
		MinIntPtr:        &i,
	}

	err := underTest.Validate()
	s.Error(err, "Valid Struct should have had an error")
	s.Require().IsType(gencheck.ValidationErrors{}, err, "Error returned was not ValidationErrors type")

	ve := err.(gencheck.ValidationErrors)
	s.Require().Len(ve, 1, "Should only have 1 validation error")

	s.Require().EqualValues(ve[0], gencheck.NewFieldError("Test", "LteTime", "lte", fmt.Errorf("is after now")), "Error should be lte time error")
}
