package generator

import (
	"fmt"
	"go/parser"
	"testing"

	"github.com/stretchr/testify/suite"
)

// BCP47TestSuite
type BCP47TestSuite struct {
	suite.Suite
}

// SetupSuite
func (s *BCP47TestSuite) SetupSuite() {
}

// TestLengthTestSuite
func TestBCP47TestSuite(t *testing.T) {
	suite.Run(t, new(BCP47TestSuite))
}

var BCP47ErrorCases = []struct {
	inType      string
	errorString string
}{
	{"int", `bcp47 is not valid on field 'TestField int'`},
	{"int8", `bcp47 is not valid on field 'TestField int8'`},
	{"int16", `bcp47 is not valid on field 'TestField int16'`},
	{"int32", `bcp47 is not valid on field 'TestField int32'`},
	{"int64", `bcp47 is not valid on field 'TestField int64'`},
	{"uint", `bcp47 is not valid on field 'TestField uint'`},
	{"uint8", `bcp47 is not valid on field 'TestField uint8'`},
	{"uint16", `bcp47 is not valid on field 'TestField uint16'`},
	{"uint32", `bcp47 is not valid on field 'TestField uint32'`},
	{"uint64", `bcp47 is not valid on field 'TestField uint64'`},
	{"bool", `bcp47 is not valid on field 'TestField bool'`},
	{"float", `bcp47 is not valid on field 'TestField float'`},
	{"float32", `bcp47 is not valid on field 'TestField float32'`},
	{"float64", `bcp47 is not valid on field 'TestField float64'`},
	{"complex64", `bcp47 is not valid on field 'TestField complex64'`},
	{"complex128", `bcp47 is not valid on field 'TestField complex128'`},
}

func (s *BCP47TestSuite) TestErrorCases() {

	format := `package test
	// SomeStruct
	type SomeStruct struct {
			TestField    %s      %s
	}`
	for _, testCase := range BCP47ErrorCases {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"bcp47\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string")

		_, err = g.Generate(f)
		s.NotNil(err, "Error for '%s' missing", testCase.inType)
		if err != nil {
			s.Contains(err.Error(), testCase.errorString)
		}
	}
}

var BCP47SuccessCases = []struct {
	inType     string
	fieldCheck string
}{
	{"*string", `gencheck.IsBCP47(s.TestField); err != nil`},
	{"string", `gencheck.IsBCP47(&s.TestField); err != nil`},
	{"[]string", `gencheck.IsBCP47(&singleTestField); err != nil`},
	{"[]*string", `gencheck.IsBCP47(singleTestField); err != nil`},
}

// TestRequiredFields will cycle through the test cases for successful calls
// to the required template and validate that the correct validation has been produced.
func (s *BCP47TestSuite) TestSuccessCases() {

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
	for _, testCase := range BCP47SuccessCases {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"bcp47\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string for type '%s'", testCase.inType)

		output, err := g.Generate(f)
		s.Nil(err, "Error generating code for input string")
		// fmt.Println(string(output))
		s.Contains(string(output), testCase.fieldCheck, "RequiredField Type='%s' ExpectedOutput='%s'", testCase.inType, testCase.fieldCheck)
	}
}
