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
	validateTag = `valid:`
)

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
}

// Field is used for storing field information.  It holds a reference to the
// original AST field information to help out if needed.
type Field struct {
	Name  string
	F     *ast.Field
	Rules []Validation
	Type  string
}

// Generator is responsible for generating validation files for the given in a go source file.
type Generator struct {
	t                     *template.Template
	knownTemplates        map[string]*template.Template
	fileSet               *token.FileSet
	generatePointerMethod bool
}

// NewGenerator is a constructor method for creating a new Generator with default
// templates loaded.
func NewGenerator() *Generator {
	g := &Generator{
		knownTemplates:        make(map[string]*template.Template),
		t:                     template.New("generator"),
		fileSet:               token.NewFileSet(),
		generatePointerMethod: false,
	}

	funcs := sprig.TxtFuncMap()

	funcs["CallTemplate"] = g.CallTemplate
	funcs["IsPtr"] = isPtr
	funcs["AddError"] = addFieldError
	funcs["IsNullable"] = isNullable
	funcs["typeof"] = typeof
	funcs["isMap"] = isMap
	funcs["isArray"] = isArray
	funcs["GenerationError"] = GenerationError
	funcs["isStruct"] = tmplIsStruct
	funcs["isStructPtr"] = tmplIsStructPtr

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

// CallTemplate is a helper method for the template to call a parsed template but with
// a dynamic name.
func (g *Generator) CallTemplate(rule Validation, data interface{}) (ret string, err error) {
	found := false
	for _, temp := range g.t.Templates() {
		if rule.Name == temp.Name() {
			found = true
			break
		}
	}
	buf := bytes.NewBuffer([]byte{})
	if !found {
		fmt.Printf("No template named for '%s' found, ignoring...\n", rule.Name)
	} else {
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

type ByFieldName []*ast.Field

func (a ByFieldName) Len() int           { return len(a) }
func (a ByFieldName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFieldName) Less(i, j int) bool { return a[i].Names[0].Name < a[j].Names[0].Name }

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
		rules := make(map[string]Field)

		fieldList := st.Fields.List

		// Go through the fields in the struct and find all the validated tags
		for _, field := range fieldList {
			// Make a field holder
			f := Field{
				F:    field,
				Name: field.Names[0].Name,
			}
			typeBuff := bytes.NewBuffer([]byte{})
			pErr := printer.Fprint(typeBuff, g.fileSet, f.F.Type)
			if pErr != nil {
				fmt.Printf("Error getting Type: %s\n", pErr)
			} else {
				f.Type = typeBuff.String()
			}

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
								// Rules are able to have parameters,
								// but will have an = in them if that is the case.

								v := Validation{
									Name:       rule,
									FieldName:  f.Name,
									F:          f.F,
									FieldType:  f.Type,
									StructName: name,
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
						}
					}

					// If we have any rules for the field, add it to the map
					if len(f.Rules) > 0 {
						rules[f.Name] = f
					}
				}
			}
		}

		data := map[string]interface{}{
			"st":        st,
			"name":      name,
			"rules":     rules,
			"ptrMethod": g.generatePointerMethod,
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
		err = fmt.Errorf("generate: error formatting code %s\n\n%s\n", err, string(vBuff.Bytes()))
	}
	return formatted, err
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
