package generator

import (
	"fmt"
	"go/parser"
	"testing"

	"github.com/stretchr/testify/suite"
)

// HexTestSuite
type HexTestSuite struct {
	suite.Suite
}

// SetupSuite
func (s *HexTestSuite) SetupSuite() {
}

// TestLengthTestSuite
func TestHexTestSuite(t *testing.T) {
	suite.Run(t, new(HexTestSuite))
}

var HexErrorCases = []struct {
	inType      string
	errorString string
}{
	{"int", `hex is not valid on field 'TestField int'`},
	{"int8", `hex is not valid on field 'TestField int8'`},
	{"int16", `hex is not valid on field 'TestField int16'`},
	{"int32", `hex is not valid on field 'TestField int32'`},
	{"int64", `hex is not valid on field 'TestField int64'`},
	{"uint", `hex is not valid on field 'TestField uint'`},
	{"uint8", `hex is not valid on field 'TestField uint8'`},
	{"uint16", `hex is not valid on field 'TestField uint16'`},
	{"uint32", `hex is not valid on field 'TestField uint32'`},
	{"uint64", `hex is not valid on field 'TestField uint64'`},
	{"bool", `hex is not valid on field 'TestField bool'`},
	{"float", `hex is not valid on field 'TestField float'`},
	{"float32", `hex is not valid on field 'TestField float32'`},
	{"float64", `hex is not valid on field 'TestField float64'`},
	{"complex64", `hex is not valid on field 'TestField complex64'`},
	{"complex128", `hex is not valid on field 'TestField complex128'`},
	{"SomeOtherStruct", `hex is not valid on field 'TestField SomeOtherStruct'`},
	{"*SomeOtherStruct", `hex is not valid on field 'TestField *SomeOtherStruct'`},
	{"SomeInterface", `hex is not valid on field 'TestField SomeInterface'`},
}

func (s *HexTestSuite) TestErrorCases() {

	format := `package test
	// SomeStruct
	type SomeStruct struct {
			TestField    %s      %s
	}`
	for _, testCase := range HexErrorCases {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"hex\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string")

		_, err = g.Generate(f)
		s.NotNil(err, "Error for '%s' missing", testCase.inType)
		if err != nil {
			s.Contains(err.Error(), testCase.errorString)
		}
	}
}

var HexSuccessCases = []struct {
	inType     string
	fieldCheck string
}{
	{"*string", `gencheck.IsHex(s.TestField); err != nil`},
	{"string", `gencheck.IsHex(&s.TestField); err != nil`},
	{"[]string", `gencheck.IsHex(&singleTestField); err != nil`},
	{"[]*string", `gencheck.IsHex(singleTestField); err != nil`},
}

// TestRequiredFields will cycle through the test cases for successful calls
// to the required template and validate that the correct validation has been produced.
func (s *HexTestSuite) TestSuccessCases() {

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
	for _, testCase := range HexSuccessCases {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"hex\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		s.Nil(err, "Error parsing input string for type '%s'", testCase.inType)

		output, err := g.Generate(f)
		s.Nil(err, "Error generating code for input string")
		// fmt.Println(string(output))
		s.Contains(string(output), testCase.fieldCheck, "RequiredField Type='%s' ExpectedOutput='%s'", testCase.inType, testCase.fieldCheck)
	}
}
