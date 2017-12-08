package generator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"sort"
	"strings"
	"text/template"

	"golang.org/x/tools/imports"

	"github.com/Masterminds/sprig"
)

const (
	validateTag  = `valid:`
	failFastFlag = `ff`
)

// Generator is responsible for generating validation files for the given in a go source file.
type Generator struct {
	t                     *template.Template
	knownTemplates        map[string]*template.Template
	fileSet               *token.FileSet
	generatePointerMethod bool
	failFast              bool
	noPrealloc            bool
}

// NewGenerator is a constructor method for creating a new Generator with default
// templates loaded.
func NewGenerator() *Generator {
	g := &Generator{
		knownTemplates:        make(map[string]*template.Template),
		t:                     template.New("generator"),
		fileSet:               token.NewFileSet(),
		generatePointerMethod: false,
		failFast:              false,
		noPrealloc:            false,
	}

	funcs := sprig.TxtFuncMap()

	funcs["CallTemplate"] = g.CallTemplate
	funcs["isPtr"] = isPtr
	funcs["addError"] = addFieldError
	funcs["isNullable"] = isNullable
	funcs["isMap"] = isMap
	funcs["isArray"] = isArray
	funcs["generationError"] = GenerationError
	funcs["isStruct"] = tmplIsStruct
	funcs["isStructPtr"] = tmplIsStructPtr
	funcs["getMapKeyType"] = g.tmplGetMapKeyType
	funcs["isParamInt"] = tmplIsParamInt

	g.t.Funcs(funcs)

	for _, assets := range AssetNames() {
		g.t = template.Must(g.t.Parse(string(MustAsset(assets))))
	}

	g.updateTemplates()

	return g
}

// WithPointerMethod is used to change the method generated for a struct to use a pointer receiver rather than a value receiver.
func (g *Generator) WithPointerMethod() *Generator {
	g.generatePointerMethod = true
	return g
}

// WithFailFast is used to change all error checks to return immediately on failure.
func (g *Generator) WithFailFast() *Generator {
	g.failFast = true
	return g
}

// WithoutPrealloc is used to remove preallocation of error array.
func (g *Generator) WithoutPrealloc() *Generator {
	g.noPrealloc = true
	return g
}

// CallTemplate is a helper method for the template to call a parsed template but with
// a dynamic name.
func (g *Generator) CallTemplate(rule Validation, data interface{}) (ret string, err error) {
	buf := bytes.NewBuffer([]byte{})

	// We don't need an else statement here because we already filter out unknown templates
	// during the parsing phase.
	if _, found := g.knownTemplates[rule.Name]; found {
		err = g.t.ExecuteTemplate(buf, rule.Name, data)
	}
	ret = buf.String()

	return
}

// GenerateFromFile is responsible for orchestrating the Code generation.  It results in a byte array
// that can be written to any file desired.  It has already had goimports run on the code before being returned.
func (g *Generator) GenerateFromFile(inputFile string) ([]byte, error) {
	f, err := g.parseFile(inputFile)
	if err != nil {
		return nil, fmt.Errorf("generate: error parsing input file '%s': %s", inputFile, err)
	}
	return g.Generate(f)

}

type byPosition []*ast.Field

func (a byPosition) Len() int           { return len(a) }
func (a byPosition) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byPosition) Less(i, j int) bool { return a[i].Pos() < a[j].Pos() }

func (g *Generator) expandEmbeddedFields(structs map[string]*ast.StructType, fields []*ast.Field) []*ast.Field {
	sort.Sort(byPosition(fields))

	actualFieldList := make([]*ast.Field, 0, len(fields))
	for _, field := range fields {
		if len(field.Names) < 1 {
			fieldName := g.getStringForExpr(field.Type)
			if embedded, ok := structs[fieldName]; ok {
				expanded := g.expandEmbeddedFields(structs, embedded.Fields.List)
				actualFieldList = append(actualFieldList, expanded...)
			}
		} else {
			actualFieldList = append(actualFieldList, field)
		}
	}

	return actualFieldList
}

// Generate does the heavy lifting for the code generation starting from the parsed AST file.
func (g *Generator) Generate(f *ast.File) ([]byte, error) {
	var err error
	structs := g.inspect(f)
	if len(structs) <= 0 {
		return nil, nil
	}

	pkg := f.Name.Name

	vBuff := bytes.NewBuffer([]byte{})
	g.t.ExecuteTemplate(vBuff, "header", map[string]interface{}{"package": pkg})

	// Make the output more consistent by iterating over sorted keys of map
	var keys []string
	for key := range structs {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, name := range keys {
		st := structs[name]

		var rules []Field

		// Only expands when we find an embedded struct
		fieldList := g.expandEmbeddedFields(structs, st.Fields.List)

		// Go through the fields in the struct and find all the validated tags
		for _, field := range fieldList {
			fieldName := ``
			// Support for embedded structs that don't have a field name associated with it
			if len(field.Names) > 0 {
				fieldName = field.Names[0].Name
			} else {
				fieldName = g.getStringForExpr(field.Type)
			}
			// Make a field holder
			f := Field{
				F:    field,
				Name: fieldName,
			}
			g.getTypeString(&f)

			if field.Tag != nil {
				if strings.Contains(field.Tag.Value, validateTag) {

					// The AST keeps the rune marker on the string, so we trim them off
					str := strings.Trim(field.Tag.Value, "`")
					// Separate tag types are separated by spaces, so split on that
					vals := strings.Split(str, " ")
					for _, val := range vals {
						// Only parse out the valid: tag
						if strings.HasPrefix(val, validateTag) {
							// Strip off the valid: prefix and the quotation marks
							ruleStr := val[len(validateTag)+1 : len(val)-1]
							// Split on commas for multiple validations
							fieldRules := strings.Split(ruleStr, ",")

							// Store the validation as the duplication check so that *if* in the future we want to support certain rules as
							// having duplicates, we can do that easier.
							dupecheck := make(map[string]Validation)
							for _, rule := range fieldRules {
								// Check for fail fast and ignore that rule
								if rule == failFastFlag {
									f.FailFast = true
									continue
								}
								// Rules are able to have parameters,
								// but will have an = in them if that is the case.

								v := Validation{
									Name:       rule,
									FieldName:  f.Name,
									F:          f.F,
									FieldType:  f.Type,
									StructName: name,
									Prealloc:   g.noPrealloc,
								}

								if strings.Contains(rule, `=`) {
									// There is a parameter, so get the rule name, and the parameter
									temp := strings.Split(rule, `=`)
									v.Name = temp[0]
									v.Param = temp[1]
								}

								// Only keep the rule if it is a known template
								if _, ok := g.knownTemplates[v.Name]; ok {
									// Make sure it's not a duplicate rule
									if _, isDupe := dupecheck[v.Name]; isDupe {
										return nil, fmt.Errorf("Duplicate rules are not allowed: '%s' on field '%s'", v.Name, f.Name)
									}
									dupecheck[v.Name] = v

									f.Rules = append(f.Rules, v)
								} else {
									fmt.Printf("Skipping unknown validation template: '%s'\n", v.Name)
								}
							}

							if f.FailFast || g.failFast {
								for index, val := range f.Rules {
									val.FailFast = true
									f.Rules[index] = val
								}
							}
						}
					}

					// If we have any rules for the field, add it to the map
					if len(f.Rules) > 0 {
						rules = append(rules, f)
					}
				}
			}
		}

		data := map[string]interface{}{
			"st":             st,
			"name":           name,
			"rules":          rules,
			"ptrMethod":      g.generatePointerMethod,
			"globalFailFast": g.failFast,
			"prealloc":       !g.noPrealloc,
		}

		if len(rules) > 0 {
			err = g.t.ExecuteTemplate(vBuff, "struct", data)

			if err != nil {
				if te, ok := err.(template.ExecError); ok {
					return nil, fmt.Errorf("generate: error executing template: %s", te.Err)
				}

				return nil, fmt.Errorf("generate: error executing template: %s", err)
			}
		}
	}

	formatted, err := imports.Process(pkg, vBuff.Bytes(), nil)
	if err != nil {
		err = fmt.Errorf("generate: error formatting code %s\n\n%s\n", err, vBuff.String())
	}
	return formatted, err
}

func (g *Generator) getTypeString(f *Field) {
	tString := g.getStringForExpr(f.F.Type)
	f.Type = tString
}

func (g *Generator) getStringForExpr(f ast.Expr) string {
	typeBuff := bytes.NewBuffer([]byte{})
	pErr := printer.Fprint(typeBuff, g.fileSet, f)
	if pErr != nil {
		fmt.Printf("Error getting Type: %s\n", pErr)
	}
	return typeBuff.String()
}

// AddTemplateFiles will be used during generation when the command line accepts
// user templates to add to the generation.
func (g *Generator) AddTemplateFiles(filenames ...string) (err error) {
	g.t, err = g.t.ParseFiles(filenames...)
	if err == nil {
		g.updateTemplates()
	}
	return
}

// updateTemplates will update the lookup map for validation checks that are
// allowed within the template engine.
func (g *Generator) updateTemplates() {
	for _, template := range g.t.Templates() {
		g.knownTemplates[template.Name()] = template
	}
}

// parseFile simply calls the go/parser ParseFile function with an empty token.FileSet
func (g *Generator) parseFile(fileName string) (*ast.File, error) {
	// Parse the file given in arguments
	return parser.ParseFile(g.fileSet, fileName, nil, parser.ParseComments)
}

// inspect will walk the ast and fill a map of names and their struct information
// for use in the generation template.
func (g *Generator) inspect(f *ast.File) map[string]*ast.StructType {

	structs := make(map[string]*ast.StructType)
	// Inspect the AST and find all structs.
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.Ident:
			if x.Obj != nil {
				// Make sure it's a Type Identifier
				if x.Obj.Kind == ast.Typ {
					// Make sure it's a spec (Type Identifiers can be throughout the code)
					if ts, ok := x.Obj.Decl.(*ast.TypeSpec); ok {
						// Only stsore the struct types (we don't do anything for interfaces)
						if sts, store := ts.Type.(*ast.StructType); store {
							structs[x.Name] = sts
						}
					}
				}
			}
		}
		// Return true to continue through the tree
		return true
	})

	return structs
}
