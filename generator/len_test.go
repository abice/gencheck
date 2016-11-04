package generator

import (
	"fmt"
	"go/parser"
	"testing"

	"github.com/stretchr/testify/suite"
)

// LenTestSuite
type LenTestSuite struct {
	suite.Suite
}

// SetupSuite
func (s *LenTestSuite) SetupSuite() {
}

// TestLengthTestSuite
func TestLenTestSuite(t *testing.T) {
	suite.Run(t, new(LenTestSuite))
}

var lenErrorCases = []struct {
	inType      string
	errorString string
}{
	{"bool", `len is not valid on field 'TestField bool'`},
	{"func()()", `len is not valid on field 'TestField func()'`},
	{"*int", `len is not valid on field 'TestField *int'`},
	{"*string", `len is not valid on field 'TestField *string'`},
	{"SomeInterface", `len is not valid on field 'TestField SomeInterface'`},
}

func (s *LenTestSuite) TestErrorCases() {

	format := `package test
	// SomeStruct
	type SomeStruct struct {
			TestField    %s      %s
	}`
	for _, testCase := range lenErrorCases {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"len=12\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string")

		_, err = g.Generate(f)
		s.Require().NotNil(err, "Error for '%s' missing", testCase.inType)
		s.Contains(err.Error(), testCase.errorString)
	}
}

var lenSuccessCases = []struct {
	inType     string
	fieldCheck string
}{
	{"int", `!(s.TestField == 12)`},
	{"int8", `!(s.TestField == 12)`},
	{"int16", `!(s.TestField == 12)`},
	{"int32", `!(s.TestField == 12)`},
	{"int64", `!(s.TestField == 12)`},
	{"uint", `!(s.TestField == 12)`},
	{"uint8", `!(s.TestField == 12)`},
	{"uint16", `!(s.TestField == 12)`},
	{"uint32", `!(s.TestField == 12)`},
	{"uint64", `!(s.TestField == 12)`},
	{"float", `!(s.TestField == 12)`},
	{"float32", `!(s.TestField == 12)`},
	{"float64", `!(s.TestField == 12)`},
	{"complex64", `!(s.TestField == 12)`},
	{"complex128", `!(s.TestField == 12)`},
	{"string", `!(len(s.TestField) == 12)`},
	{"chan int", `!(len(s.TestField) == 12)`},
	{"[]string", `!(len(s.TestField) == 12)`},
	{"map[string]string", `!(len(s.TestField) == 12)`},
	{"[]SomeOtherStruct", `!(len(s.TestField) == 12)`},
	{"[]SomeInterface", `!(len(s.TestField) == 12)`},
}

// TestRequiredFields will cycle through the test cases for successful calls
// to the required template and validate that the correct validation has been produced.
func (s *LenTestSuite) TestSuccessCases() {

	format := `package test
	// SomeInterface
	type SomeInterface interface{

	}
	// SomeOtherStruct
	type SomeOtherStruct struct{

	}
	// SomeStruct
	type SomeStruct struct {
			TestField    %s      %s
	}`
	for _, testCase := range lenSuccessCases {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"len=12\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string for type '%s'", testCase.inType)

		output, err := g.Generate(f)
		s.Nil(err, "Error generating code for input string")
		// fmt.Println(string(output))
		s.Contains(string(output), testCase.fieldCheck, "RequiredField Type='%s' ExpectedOutput='%s'", testCase.inType, testCase.fieldCheck)
	}
}
