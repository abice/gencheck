package generator

import (
	"fmt"
	"go/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

var otherPkgDiveTests = []struct {
	inType     string
	fieldCheck string
}{
	{"example.Inner", `gencheck.Validate(s.TestField)`},
	{"*example.Inner", `gencheck.Validate(s.TestField)`},
	{"example.IFace", `gencheck.Validate(s.TestField)`},
}

func TestIsFuncSelector(t *testing.T) {
	format := `package test

	// SomeStruct
	type SomeStruct struct {
			TestField    %s      %s
	}`
	for _, testCase := range otherPkgDiveTests {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, "`valid:\"dive\"`")

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		assert.Nil(t, err, "Error parsing input string for type '%s'", testCase.inType)

		output, err := g.Generate(f)
		assert.Nil(t, err, "Error generating code for input string")
		// fmt.Println(string(output))
		assert.Contains(t, string(output), testCase.fieldCheck, "RequiredField Type='%s' ExpectedOutput='%s'", testCase.inType, testCase.fieldCheck)
	}
}

var mapContainsTests = []struct {
	inType     string
	validTag   string
	fieldCheck string
}{
	{"map[string]interface{}", "`valid:\"contains=myKey\"`", `s.MapOfSomething["myKey"]; !foundMapOfSomething`},
	{"map[int]interface{}", "`valid:\"contains=10\"`", `s.MapOfSomething[10]; !foundMapOfSomething`},
}

func TestMapContains(t *testing.T) {
	format := `package test

	// SomeStruct
	type SomeStruct struct {
			MapOfSomething    %s      %s
	}`
	for _, testCase := range mapContainsTests {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, testCase.validTag)

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		assert.Nil(t, err, "Error parsing input string for type '%s'", testCase.inType)

		output, err := g.Generate(f)
		assert.Nil(t, err, "Error generating code for input string")
		// fmt.Println(string(output))
		assert.Contains(t, string(output), testCase.fieldCheck, "RequiredField Type='%s' ExpectedOutput='%s'", testCase.inType, testCase.fieldCheck)
	}
}

var failedMapContainsTests = []struct {
	inType      string
	validTag    string
	errorString string
}{
	{"map[int]interface{}", "`valid:\"contains=myKey\"`", `cannot use a non integer value on an integer keyed map`},
}

func TestFailedMapContains(t *testing.T) {
	format := `package test

	// SomeStruct
	type SomeStruct struct {
			MapOfSomething    %s      %s
	}`
	for _, testCase := range failedMapContainsTests {

		g := NewGenerator()
		input := fmt.Sprintf(format, testCase.inType, testCase.validTag)

		f, err := parser.ParseFile(g.fileSet, "TestRequiredErrors", input, parser.ParseComments)
		assert.Nil(t, err, "Error parsing input string for type '%s'", testCase.inType)

		_, err = g.Generate(f)
		assert.NotNil(t, err, "Error on required int missing")
		assert.Contains(t, err.Error(), testCase.errorString)
	}
}
