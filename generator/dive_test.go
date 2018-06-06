package generator

import (
	"fmt"
	"go/parser"
	"testing"

	"github.com/stretchr/testify/suite"
)

// DiveTestSuite
type DiveTestSuite struct {
	suite.Suite
}

// SetupSuite
func (s *DiveTestSuite) SetupSuite() {
}

// TestDiveTestSuite
func TestDiveTestSuite(t *testing.T) {
	suite.Run(t, new(DiveTestSuite))
}

var diveErrorCases = []struct {
	inType      string
	errorString string
}{
	{"bool", `Dive is not valid on field 'TestField bool'`},
	{"*int", `Dive is not valid on field 'TestField *int'`},
	//{"[]*string", `Dive is not valid on field 'TestField []*string'`},		//dive does not reject validating slice elements missing validators
	//{"map[string]int", `Dive is not valid on field 'TestField map[string]int'`},		//dive does not reject validating map elements missing validators
}

func (s *DiveTestSuite) TestErrorCases() {

	format := `package test
	// SomeStruct
	type SomeStruct struct {
			TestField    %s      %s
	}`
	for _, testCase := range diveErrorCases {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"dive\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string")

		_, err = g.Generate(f)
		s.Require().NotNil(err, "Error for '%s' missing", testCase.inType)
		s.Contains(err.Error(), testCase.errorString)
	}
}

var diveSuccessCases = []struct {
	inType     string
	fieldCheck string
}{
	{"SomeOtherStruct", `gencheck.Validate(s.TestField)`},
	{"*SomeOtherStruct", `gencheck.Validate(s.TestField)`},
	{"[]SomeOtherStruct", `gencheck.Validate(e)`},
	{"[]*SomeOtherStruct", `gencheck.Validate(e)`},
	{"map[string]SomeOtherStruct", `gencheck.Validate(e)`},
	{"map[string]*SomeOtherStruct", `gencheck.Validate(e)`},
}

func (s *DiveTestSuite) TestSuccessCases() {

	format := `package test
	// SomeOtherStruct
	type SomeOtherStruct struct{
		RequiredField 	string 	%s
	}
	// SomeStruct
	type SomeStruct struct {
			TestField    %s      %s
	}`
	for _, testCase := range diveSuccessCases {

		g := NewGenerator()
		input := fmt.Sprintf(format, "`valid:\"required\"`", testCase.inType, "`valid:\"dive\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string for type '%s'", testCase.inType)

		output, err := g.Generate(f)
		s.Nil(err, "Error generating code for input string")
		s.Contains(string(output), testCase.fieldCheck, "RequiredField Type='%s' ExpectedOutput='%s'", testCase.inType, testCase.fieldCheck)
	}
}
