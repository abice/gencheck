# gencheck
[![CircleCI](https://circleci.com/gh/abice/gencheck.svg?style=svg&circle-token=db146e1d9c8d935d7bd05ac879f818801c432ea4)](https://circleci.com/gh/abice/gencheck)
[![Coverage Status](https://coveralls.io/repos/github/abice/gencheck/badge.svg)](https://coveralls.io/github/abice/gencheck)
[![Go Report Card](https://goreportcard.com/badge/github.com/abice/gencheck)](https://goreportcard.com/report/github.com/abice/gencheck)

### Validation generation for go.

### Built-in validations
See [Validations](validations.md)

## How it works
gencheck was built using the idea of [zencoder/gokay](https://github.com/zencoder/gokay), but uses templates to create validations for a struct.

gencheck will use the `valid` tag within a struct to generate a `Validate()` method, which is will store in a `file_validators.go` file
next to the input file.

gencheck's `Validate()` method will return a `ValidationErrors` type, which is an array of `FieldError`s.

Given the struct:

```go
type MyStruct struct{
	MyField string `valid:"required"`
}
```
A Validate method is generated:

```go
func (s MyStruct) Validate() error {
	var vErrors gencheck.ValidationErrors

	// BEGIN MyField Validations
	// required
	if s.MyField == "" {
		vErrors = append(vErrors, gencheck.NewFieldError("MyStruct", "MyField", "required", errors.New("is required")))
	}
	// END MyField Validations

	if len(vErrors) > 0 {
		return vErrors
	}
	return nil
}
```

## Installing

First use `go get` to install the latest version of the library.

`go get -v github.com/abice/gencheck/gencheck`

Normally the above command will build and install the binary, but just to be sure it is installed in your `GOPATH` for use on the command line, or build args:

`go install github.com/abice/gencheck/gencheck`

## Running
### Command line
```sh
gencheck -f=file.go -t="SomeTemplate.tmpl" --template="SomeOtherTemplate.tmpl" -d="some/dir" --template-dir="some/dir/that/has/templates"
```

### Using with `go generate`
Add a `//go:generate` tag to the top of your file that you want to generate for, including the file name.

```go
//go:generate gencheck -f=this_file.go
```

## Adding Validations
Add validations to `valid` tag in struct def:

```go
type ExampleStruct struct {
	HexStringPtr            *string `valid:"len=16,notnil,hex"`
	HexString               string  `valid:"len=12,hex"`
	CanBeNilWithConstraints *string `valid:"len=12"`
}
```


### Tag syntax
Validation tags are comma separated, with any validation parameter specified after an equal sign.

`valid:"ValidationName1,ValidationName2=vn2param"`

In the above example, the `hex` and `notnil` Validations are parameterless, whereas len requires 1 parameter.


### Time comparisons
Since the addition of `gt(e)` and `lt(e)`, there are now comparisons for `time.Time` values.  If no arguments are specified to those, then it calculates whether the
field time is After and Before `time.Now().UTC()` respectively.  You can specify a parameter for those validations if you choose.  The parameter will be interpreted as
the offset to use with respect to `time.Now().UTC()` by utilizing the `Add()` function.

```go
requestTime time.Time `valid:"gte=-1*time.Second"`
```

```go
tGteTimeVal := time.Now().UTC().Add(-1 * time.Second)
if s.GteTimeVal.Before(tGteTimeVal) {
	vErrors = append(vErrors, gencheck.NewFieldError("Test", "GteTimeVal", "gte", fmt.Errorf("is before %s", tGteTimeVal)))
}
```

### Fail Fast flag
The fail fast flag is a built-in validation flag that will allow you to return immediately on an invalid check.  This allows you to not waste time checking the rest of the struct if a vital field is wrong.  It can be placed anywhere within the `valid` tag, and will be applied to **all** rules within that field.

There is also a `--failfast` flag on the cli that will allow you to make **all** validations within **all** structs found in the files to be fail fast.

### Writing your own Validations
gencheck allows developers to write and attach their own Validation templates to the generator.

1. Write a template that creates a validation for a given field making sure to define the template as the validation tag you want to use:
	
	```gotemplate
	{{define "mycheck" -}}
	if err := gencheck.IsUUID({{.Param}}, {{if not (IsPtr . )}}&{{end}}s.{{.FieldName}}); err != nil {
	  {{ AddError . "err" }}
	}
	{{end -}}
	```

1. Import that template when running gencheck
1. Write tests for your struct's constraints
1. Add `valid` tags to your struct fields
1. Run gencheck: `gencheck -f=file.go -t=MyTemplate`

NOTES:

- In your template, the . pipeline is an instance of the [`generator.Validation`](generator/types.go#L27) struct.
- The template functions from [Sprig](https://github.com/Masterminds/sprig) have been included.
- There are some custom functions provided for you to help in determining the ast field type
 - isPtr
 - addError
 - isNullable
 - isMap
 - isArray
 - isStruct
 - isStructPtr
 - isStructPtr
 - generationError
   - Allows you to fail code generation with a specific error message

[More Examples](internal/example/)


## Useless Benchmarks

I know benchmarks are always skewed to show what the creators want you to see, but here's a quick benchmark of the cost of using validation to check.

I've also added some comparison benchmark output from the `./internal/benchmark_test.go` to compare the different options with gencheck and how it holds up to the go playground validator.

 - [Playground](benchmark_playground.md)
 - [No Options](benchmark_nooptions.md)
 - [No Preallocated error array](benchmark_noprealloc.md)
 - [Fail Fast](benchmark_failfast.md)

```
BenchmarkReflectionInt-8      	20000000	       104 ns/op
BenchmarkEmptyInt-8           	2000000000	         0.29 ns/op
BenchmarkReflectionStruct-8   	 5000000	       262 ns/op
BenchmarkEmptyStruct-8        	50000000	        28.3 ns/op
BenchmarkReflectionString-8   	10000000	       159 ns/op
BenchmarkEmptyString-8        	200000000	         9.49 ns/op

```
Benchmarks using fail fast flag
```
BenchmarkValidString-8            	300000000	         5.02 ns/op
BenchmarkFailing1TestString-8     	10000000	       158 ns/op
BenchmarkFailing2TestString-8     	10000000	       159 ns/op
BenchmarkFailingAllTestString-8   	10000000	       164 ns/op
```

Benchmarks without fail fast flag and preallocated capacity for errors
```
BenchmarkValidString-8            	20000000	        68.7 ns/op
BenchmarkFailing1TestString-8     	10000000	       189 ns/op
BenchmarkFailing2TestString-8     	 5000000	       272 ns/op
BenchmarkFailingAllTestString-8   	 3000000	       418 ns/op
```

## Development

### Dependencies

Tested on go 1.7.3.

### Build and run unit tests

    make test

### TODO
- [x] Testing for templates
- [x] Prevent duplicate validations on the same field
- [x] Update Required tag to error out on numerical or boolean fields
- [ ] Support for sub-validations? `Struct fields: generated code will call static Validate method on any field that implements Validateable interface`  Maybe use a deep check
- [x] Readme info for what information is available within the templates.
- [ ] Contains for other slice types.
- [ ] Contains for maps.
- [ ] Add support for build tags for generated file.
- [ ] Cross field validation (i.e. x.start <= x.end)

### CI

[This library builds on Circle CI, here.](https://circleci.com/gh/abice/gencheck/)

## License

[Apache License Version 2.0](LICENSE)
