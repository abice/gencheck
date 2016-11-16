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
