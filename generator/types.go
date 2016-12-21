package generator

import "go/ast"

// Validation is a holder for a validation rule within the generation templates
// It actually has the same information as the Field struct, simply for ease of
// access from within the templates.
type Validation struct {
	Name       string
	Param      string
	FieldName  string
	F          *ast.Field
	FieldType  string
	StructName string
	FailFast   bool
	Prealloc   bool
}

// Field is used for storing field information.  It holds a reference to the
// original AST field information to help out if needed.
type Field struct {
	Name     string
	F        *ast.Field
	Rules    []Validation
	Type     string
	FailFast bool
}
