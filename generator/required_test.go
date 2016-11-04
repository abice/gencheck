package generator

import (
	"fmt"
	"go/parser"
	"testing"

	"github.com/stretchr/testify/suite"
)

// RequiredTestSuite
type RequiredTestSuite struct {
	suite.Suite
}

// SetupSuite
func (s *RequiredTestSuite) SetupSuite() {
}

// TestLengthTestSuite
func TestRequiredTestSuite(t *testing.T) {
	suite.Run(t, new(RequiredTestSuite))
}

var requiredErroringTests = []struct {
	inType      string
	errorString string
}{
	{"int", `Required on integer values is not supported.`},
	{"int8", `Required on integer values is not supported.`},
	{"int16", `Required on integer values is not supported.`},
	{"int32", `Required on integer values is not supported.`},
	{"int64", `Required on integer values is not supported.`},
	{"uint", `Required on integer values is not supported.`},
	{"uint8", `Required on integer values is not supported.`},
	{"uint16", `Required on integer values is not supported.`},
	{"uint32", `Required on integer values is not supported.`},
	{"uint64", `Required on integer values is not supported.`},
	{"bool", `Required on boolean values is not supported.`},
	{"float", `Required on numerical values is not supported.`},
	{"float32", `Required on numerical values is not supported.`},
	{"float64", `Required on numerical values is not supported.`},
	{"complex64", `Required on numerical values is not supported.`},
	{"complex128", `Required on numerical values is not supported.`},
}

func (s *RequiredTestSuite) TestRequiredErrors() {

	format := `package test
	// SomeStruct
	type SomeStruct struct {
			Test%s    %s      %s
	}`
	for _, testCase := range requiredErroringTests {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, testCase.inType, "`valid:\"required\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string")

		_, err = g.Generate(f)
		s.NotNil(err, "Error on required int missing")
		s.Contains(err.Error(), testCase.errorString)
	}
}

var requiredSuccessTests = []struct {
	inType     string
	fieldCheck string
}{
	{"chan int", `s.TestField == nil`},
	{"func()()", `s.TestField == nil`},
	{"*int", `s.TestField == nil`},
	{"*string", `s.TestField == nil`},
	{"SomeInterface", `s.TestField == nil`},
	{"[]string", `s.TestField == nil`},
	{"map[string]string", `s.TestField == nil`},
	{"string", `s.TestField == ""`},
	{"SomeOtherStruct", `s.TestField == zeroTestField`},
}

// TestRequiredFields will cycle through the test cases for successful calls
// to the required template and validate that the correct validation has been produced.
func (s *RequiredTestSuite) TestRequiredFields() {

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
	for _, testCase := range requiredSuccessTests {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"required\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string for type '%s'", testCase.inType)

		output, err := g.Generate(f)
		s.Nil(err, "Error generating code for input string")
		// fmt.Println(string(output))
		s.Contains(string(output), testCase.fieldCheck, "RequiredField Type='%s' ExpectedOutput='%s'", testCase.inType, testCase.fieldCheck)
	}
}
