package example

import (
	"fmt"
	"testing"

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
func (s *ExampleTestSuite) TestValidateTestStruct_NoValues() {
	expected := gencheck.ValidationErrors{
		gencheck.NewFieldError("Test", "RequiredString", "required", fmt.Errorf("is required")),
		gencheck.NewFieldError("Test", "RequiredMultiple", "required", fmt.Errorf("is required")),
		gencheck.NewFieldError("Test", "LenString", "len", fmt.Errorf("length mismatch")),
		gencheck.NewFieldError("Test", "LenNumber", "len", fmt.Errorf("length mismatch")),
		gencheck.NewFieldError("Test", "LenMultiple", "len", fmt.Errorf("length mismatch")),
		gencheck.NewFieldError("Test", "UUID", "uuid", fmt.Errorf("'' is not a UUID")),
	}

	underTest := Test{}

	err := underTest.Validate()
	s.Require().IsType(gencheck.ValidationErrors{}, err, "Error returned was not ValidationErrors type")

	ve := err.(gencheck.ValidationErrors)

	s.Require().Equal(len(expected), len(ve), "Validation Errors were of different length")

	for i, fe := range ve {
		s.Equal(expected[i], fe)
	}
}

// TestValidateTestStruct_NoValues
func (s *ExampleTestSuite) TestValidateTestStruct_Values() {

	underTest := Test{
		LenMultiple:      []string{"", "", "", "", "", "", ""},
		LenNumber:        1113,
		LenString:        "a",
		RequiredMultiple: []string{},
		RequiredString:   "b",
		UUID:             "7112EE37-3219-4A26-BA01-1D230BC9257B",
	}

	err := underTest.Validate()
	s.Nil(err, "Valid Struct should not have had an error")

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
