# gencheck [![CircleCI](https://circleci.com/gh/zencoder/gokay.svg?style=svg&circle-token=90f42bc5cbb6fe74834f7649d67298130431d88d)](https://circleci.com/gh/zencoder/gokay) [![Coverage Status](https://coveralls.io/repos/github/zencoder/gokay/badge.svg?branch=circle-fixes&t=A2kWWv)](https://coveralls.io/github/zencoder/gokay?branch=circle-fixes)
Validation generator for go.

## How it works
gencheck will use the `valid` tag within a struct to generate a `Validate()` method.

gencheck's `Validate()` method will return a `ValidationErrors` type, which is an array of `FieldError`s.

### Code Documentation
`godoc -http=:6060`

## Installing

First use `go get` to install the latest version of the library.

`go get -v github.com/abice/gencheck/cmd`

Then install the binary in your `GOPATH` for use on the command line, or build args.

`go install github.com/abice/gencheck/cmd`

## Running
### Usage
```	sh
gencheck -f file.go -t="SomeTemplate.tmpl" --template="SomeOtherTemplate.tmpl" -d="some/dir" --template-dir="some/dir/that/has/templates"
```

### Examples
Generate validation methods for types in a single file
```sh
gencheck -f file.go
```

Generate validation methods with a custom package and constructor
```sh
gokay file.go --template=CustomGenerator.tmpl
```

gokay relies on the goimports tool to format and resolve imports of the generated validation file.

## Using gokay
- Add validations to `valid` tag in struct def:

```go
type ExampleStruct struct {
	HexStringPtr            *string `valid:"Length=16,NotNil,Hex"`
	HexString               string  `valid:"Length=12,Hex"`
	CanBeNilWithConstraints *string `valid:"Length=12"`
}
```

- Run gokay command

### Tag syntax
Validation tags are comma separated, with any validation parameter specified after an equal sign.

`valid:"ValidationName1,ValidationName2=vn2param"`

In the above example, the `Hex` and `NotNil` Validations are parameterless, whereas length requires 1 parameter.

### Built-in Validations
Name | Params | Allowed Field Types | Description
---- | ------------------- | ------ | -----------
Hex  | N/A | `(*)string` | Checks if a string field is a valid hexadecimal format number (0x prefix optional)
NotNil | N/A | pointers | Checks and fails if a pointer is nil
Length | 1 | `(*)string` | Checks if a string's length matches the tag's parameter
UUID | N/A | `(*)string` | Checks and fails if a string is not a valid UUID

### Implicitly generated validations
These sections of code will be added to the generated `Validate()` method regardless of a field's `valid` tag's contents.
If a struct does not have any `valid` tags and no fields with implicit validation rules, then no Validate method will be generated.

- Struct fields: generated code will call static Validate method on any field that implements Validateable interface
- Slice/Map fields: Static Validate will be called on each element of a slice or map of structs or struct pointers (one level of indirection). Only supports maps with string indices.


*Note* on built-in and implicit validations: With the obvious exception of NotNil, nil pointers fields are considered to be valid in order to allow support for optional fields.

### Writing your own Validations
gokay was built to allow developers to write and attach their own Validations to the Validate generator.

1. Write a template that creates a validation for a given field making sure to define the template as the validation tag you want to use:

    ```go
		{{define "UUID"}}
		if err := gokay.IsUUID({{.Param}}, {{if not (IsPtr . )}}&{{end}}s.{{.FieldName}}); err != nil {
		  {{ AddError . "err" }}
		}
		{{end -}}
    ```

1. Import that template when running gokay
1. Write tests for your struct's constraints
1. Add `valid` tags to your struct fields
1. Run gokay: `gokay file.go --template=MyTemplate`

[More Examples](internal/gkexample/)

## Development

### Dependencies

Tested on go 1.7.1.

### Build and run unit tests

    make test

### TODO
- [ ] Testing for templates
- [x] Prevent duplicate validations on the same field
- [x] Update Required tag to error out on numerical or boolean fields
- [ ] Support for sub-validations? `Struct fields: generated code will call static Validate method on any field that implements Validateable interface`  Maybe use a deep check
- [ ] Move cli to cmd directory, so that the gokay pkg is at the root of the repo.

### CI

[This library builds on Circle CI, here.](https://circleci.com/gh/zencoder/gokay/)

## License

[Apache License Version 2.0](LICENSE)
