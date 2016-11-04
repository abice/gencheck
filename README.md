# gencheck [![CircleCI](https://circleci.com/gh/zencoder/gencheck.svg?style=svg&circle-token=90f42bc5cbb6fe74834f7649d67298130431d88d)](https://circleci.com/gh/zencoder/gencheck) [![Coverage Status](https://coveralls.io/repos/github/zencoder/gencheck/badge.svg?branch=circle-fixes&t=A2kWWv)](https://coveralls.io/github/zencoder/gencheck?branch=circle-fixes)
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

## Adding Validations
- Add validations to `valid` tag in struct def:

```go
type ExampleStruct struct {
	HexStringPtr            *string `valid:"len=16,notnil,hex"`
	HexString               string  `valid:"len=12,hex"`
	CanBeNilWithConstraints *string `valid:"len=12"`
}
```

- Run gencheck

### Tag syntax
Validation tags are comma separated, with any validation parameter specified after an equal sign.

`valid:"ValidationName1,ValidationName2=vn2param"`

In the above example, the `hex` and `notnil` Validations are parameterless, whereas len requires 1 parameter.

### Built-in Validations
Name | Params | Allowed Field Types | Description
---- | ------------------- | ------ | -----------
hex  | N/A | `(*)string` | Checks if a string field is a valid hexadecimal format number (0x prefix optional)
notnil | N/A | pointers | Checks and fails if a pointer is nil
len | 1 | `(*)string` | Checks if a string's length matches the tag's parameter
uuid | N/A | `(*)string` | Checks and fails if a string is not a valid UUID

### Writing your own Validations
gencheck allows developers to write and attach their own Validation templates to the generator.

1. Write a template that creates a validation for a given field making sure to define the template as the validation tag you want to use:

    ```go
		{{define "mycheck"}}
		if err := gencheck.IsUUID({{.Param}}, {{if not (IsPtr . )}}&{{end}}s.{{.FieldName}}); err != nil {
		  {{ AddError . "err" }}
		}
		{{end -}}
    ```

1. Import that template when running gencheck
1. Write tests for your struct's constraints
1. Add `valid` tags to your struct fields
1. Run gencheck: `gencheck -f=file.go -t=MyTemplate`

[More Examples](internal/gkexample/)

## Development

### Dependencies

Tested on go 1.7.3.

### Build and run unit tests

    make test

### TODO
- [ ] Testing for templates
- [x] Prevent duplicate validations on the same field
- [x] Update Required tag to error out on numerical or boolean fields
- [ ] Support for sub-validations? `Struct fields: generated code will call static Validate method on any field that implements Validateable interface`  Maybe use a deep check

### CI

[This library builds on Circle CI, here.](https://circleci.com/gh/abice/gencheck/)

## License

[Apache License Version 2.0](LICENSE)
