package generator

import (
	"errors"
	"fmt"
	"go/ast"
	"strconv"
	"strings"
)

const (
	errorFormat                = `vErrors = append(vErrors, gencheck.NewFieldError("%s","%s","%s",%s))`
	failFastErrorFormat        = `return append(vErrors, gencheck.NewFieldError("%s","%s","%s",%s))`
	indexedErrorFormat         = `vErrors = append(vErrors, gencheck.NewFieldError("%s",fmt.Sprintf("%s[%%v]", %s),"%s",%s))`
	indexedFailFastErrorFormat = `return append(vErrors, gencheck.NewFieldError("%s",fmt.Sprintf("%s[%%v]", %s),"%s",%s))`
)

// addFieldError is a helper method for templates to add an error to a field.
func addFieldError(v Validation, eString string) (ret string, err error) {
	if v.FailFast {
		ret = fmt.Sprintf(failFastErrorFormat, v.StructName, v.FieldName, v.Name, eString)
	} else {
		ret = fmt.Sprintf(errorFormat, v.StructName, v.FieldName, v.Name, eString)
	}
	return
}

// addIndexedFieldError is a helper method for templates to add an indexed error to a field.
func addIndexedFieldError(v Validation, fieldIndex, eString string) (ret string, err error) {
	if v.FailFast {
		ret = fmt.Sprintf(indexedFailFastErrorFormat, v.StructName, v.FieldName, fieldIndex, v.Name, eString)
	} else {
		ret = fmt.Sprintf(indexedErrorFormat, v.StructName, v.FieldName, fieldIndex, v.Name, eString)
	}
	return
}

// accessor is a helper method for templates to use to remove a lot of `if isPtr *`
func accessor(f Validation, prefix string) (string, error) {
	isPtr, err := isPtr(f)
	if err != nil || !isPtr {
		return fmt.Sprintf("%s%s", prefix, f.FieldName), nil
	}
	return fmt.Sprintf("*%s%s", prefix, f.FieldName), nil
}

// IsPtr is a helper method for templates to use to determine if a field is a pointer.
func isPtr(f Validation) (ret bool, err error) {
	ret = false
	field := f.F
	// fmt.Printf("ast.Field: %#v\n", field.Type)
	_, ret = field.Type.(*ast.StarExpr)
	return
}

// IsNullable is a helper method for templates to use to determine if a field is a nullable
func isNullable(f Validation) (ret bool, err error) {
	ret = false
	field := f.F
	fType := field.Type
	if _, ok := fType.(*ast.Ident); ok {
		fType = getIdentType(fType.(*ast.Ident))
	}

	// fmt.Printf("IsNullable: %s=>%#v\n", f.FieldType, fType)
	switch fType.(type) {
	case *ast.StarExpr:
		ret = true
	case *ast.MapType:
		ret = true
	case *ast.ArrayType:
		ret = true
	case *ast.ChanType:
		ret = true
	case *ast.FuncType:
		ret = true
	case *ast.InterfaceType:
		ret = true
	}
	return
}

// isArray is a helper method for templates to determine if the field is an array
func isArray(f Validation) (ret bool, err error) {
	ret = false
	field := f.F
	switch field.Type.(type) {
	case *ast.ArrayType:
		ret = true
	}
	return
}

// isMap is a helper method for templates to determine if the field is a map
func isMap(f Validation) (ret bool, err error) {
	ret = false
	field := f.F
	switch field.Type.(type) {
	case *ast.MapType:
		ret = true
	}
	return
}

// GenerationError is a helper method for templates to generate an error if something
// is unsupported.
func GenerationError(s string) (ret interface{}, err error) {
	err = errors.New(s)
	return
}

// isStruct is a helper method for templates to determine if the field is a struct (not a pointer to a struct)
func (g *Generator) tmplIsStruct(f Validation) (ret bool, err error) {
	ret = g.isStruct(f.F)
	return
}

// isStruct is a helper method for templates to determine if the field is a struct (not a pointer to a struct)
func (g *Generator) isStruct(field *ast.Field) (ret bool) {
	ret = false
	fType := field.Type
	typeString := g.getStringForExpr(field.Type)
	if strings.HasPrefix(typeString, `*`) {
		// Quit early if it's a pointer
		return
	}
	if _, knownStruct := g.knownStructs[typeString]; knownStruct {
		// We know it's not a pointer, so if we can find it, it's a struct
		ret = true
		return
	}

	if _, ok := fType.(*ast.Ident); ok {
		fType = getIdentType(fType.(*ast.Ident))
	}

	switch fType.(type) {
	case *ast.StructType:
		ret = true
	case *ast.SelectorExpr:
		ret = true
	}
	return
}

// tmplIsStructPtr is the template helper method to determine if a field is a *Struct
func (g *Generator) tmplIsStructPtr(f Validation) (ret bool, err error) {
	ret = g.isStructPtr(f.F)
	return
}

// isStructPtr is a helper method for templates to determine if the field is a pointer to a struct
func (g *Generator) isStructPtr(field *ast.Field) (ret bool) {
	ret = false
	fType := field.Type
	typeString := g.getStringForExpr(field.Type)
	if !strings.HasPrefix(typeString, `*`) {
		// Quit early if it's not a pointer at all
		return
	}
	if _, knownStruct := g.knownStructs[strings.TrimPrefix(typeString, `*`)]; knownStruct {
		// We already know it's a pointer, so if it's found, return true
		ret = true
		return
	}

	if star, ok := fType.(*ast.StarExpr); ok {
		switch star.X.(type) {
		case *ast.Ident:
			fType = getIdentType(star.X.(*ast.Ident))
		case *ast.SelectorExpr:
			ret = true
			// fType = getIdentType(star.X.(*ast.SelectorExpr).Sel)
			// fmt.Printf("Got a selector, inner type is: %#v\n", fType)
		}
	}

	switch fType.(type) {
	case *ast.StructType:
		ret = true
	}
	return
}

func getIdentType(ident *ast.Ident) ast.Expr {
	if ident.Obj != nil {
		switch ident.Obj.Decl.(type) {
		case *ast.TypeSpec:
			return ident.Obj.Decl.(*ast.TypeSpec).Type
		}
	}
	return ident
}

// tmplIsStructPtr is the template helper method to determine if a field is a *Struct
func (g *Generator) tmplGetMapKeyType(f Validation) (ret string, err error) {
	expr := g.getMapKeyType(f.F)
	if expr != nil {
		ret = g.getStringForExpr(expr)
	}
	return
}

func (g *Generator) getMapKeyType(field *ast.Field) ast.Expr {
	switch field.Type.(type) {
	case *ast.MapType:
		mt := field.Type.(*ast.MapType)

		switch mt.Key.(type) {
		case *ast.Ident:
			return getIdentType(mt.Key.(*ast.Ident))
		}
	}
	return nil
}

func tmplIsParamInt(param string) bool {
	_, err := strconv.Atoi(param)
	if err != nil {
		return false
	}
	return true
}
