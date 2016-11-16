package generator

import (
	"errors"
	"fmt"
	"go/ast"
)

const (
	errorFormat         = `vErrors = append(vErrors, gencheck.NewFieldError("%s","%s","%s",%s))`
	failFastErrorFormat = `return append(vErrors, gencheck.NewFieldError("%s","%s","%s",%s))`
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
func tmplIsStruct(f Validation) (ret bool, err error) {
	ret = isStruct(f.F)
	return
}

// isStruct is a helper method for templates to determine if the field is a struct (not a pointer to a struct)
func isStruct(field *ast.Field) (ret bool) {
	ret = false
	fType := field.Type
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
func tmplIsStructPtr(f Validation) (ret bool, err error) {
	ret = isStructPtr(f.F)
	return
}

// isStructPtr is a helper method for templates to determine if the field is a pointer to a struct
func isStructPtr(field *ast.Field) (ret bool) {
	ret = false
	fType := field.Type
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
