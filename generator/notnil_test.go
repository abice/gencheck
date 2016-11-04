package generator

import (
	"fmt"
	"go/parser"
	"testing"

	"github.com/stretchr/testify/suite"
)

// notnilTestSuite
type notnilTestSuite struct {
	suite.Suite
}

// SetupSuite
func (s *notnilTestSuite) SetupSuite() {
}

// TestLengthTestSuite
func TestnotnilTestSuite(t *testing.T) {
	suite.Run(t, new(notnilTestSuite))
}

var notnilErrorCases = []struct {
	inType      string
	errorString string
}{
	{"int", `notnil is not valid on non nullable field 'TestField int'`},
	{"int8", `notnil is not valid on non nullable field 'TestField int8'`},
	{"int16", `notnil is not valid on non nullable field 'TestField int16'`},
	{"int32", `notnil is not valid on non nullable field 'TestField int32'`},
	{"int64", `notnil is not valid on non nullable field 'TestField int64'`},
	{"uint", `notnil is not valid on non nullable field 'TestField uint'`},
	{"uint8", `notnil is not valid on non nullable field 'TestField uint8'`},
	{"uint16", `notnil is not valid on non nullable field 'TestField uint16'`},
	{"uint32", `notnil is not valid on non nullable field 'TestField uint32'`},
	{"uint64", `notnil is not valid on non nullable field 'TestField uint64'`},
	{"bool", `notnil is not valid on non nullable field 'TestField bool'`},
	{"float", `notnil is not valid on non nullable field 'TestField float'`},
	{"float32", `notnil is not valid on non nullable field 'TestField float32'`},
	{"float64", `notnil is not valid on non nullable field 'TestField float64'`},
	{"complex64", `notnil is not valid on non nullable field 'TestField complex64'`},
	{"complex128", `notnil is not valid on non nullable field 'TestField complex128'`},
	{"string", `notnil is not valid on non nullable field 'TestField string'`},
}

func (s *notnilTestSuite) TestErrorCases() {

	format := `package test
	// SomeStruct
	type SomeStruct struct {
			TestField    %s      %s
	}`
	for _, testCase := range notnilErrorCases {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"notnil\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string")

		_, err = g.Generate(f)
		s.NotNil(err, "Error on required int missing")
		if err != nil {
			s.Contains(err.Error(), testCase.errorString)
		}
	}
}

var notnilSuccessCases = []struct {
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
	{"*SomeOtherStruct", `s.TestField == nil`},
}

// TestRequiredFields will cycle through the test cases for successful calls
// to the required template and validate that the correct validation has been produced.
func (s *notnilTestSuite) TestSuccessCases() {

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
	for _, testCase := range notnilSuccessCases {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"notnil\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string for type '%s'", testCase.inType)

		output, err := g.Generate(f)
		s.Nil(err, "Error generating code for input string")
		// fmt.Println(string(output))
		s.Contains(string(output), testCase.fieldCheck, "RequiredField Type='%s' ExpectedOutput='%s'", testCase.inType, testCase.fieldCheck)
	}
}
