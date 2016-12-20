package benchmark

import (
	"testing"

	"github.com/abice/gencheck"
	uuid "github.com/satori/go.uuid"
	"gopkg.in/go-playground/validator.v9"
)

var result bool
var errResult error
var EmptyStructs = []TestString{
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
	{},
}

// BenchmarkEmptyInt is a quick benchmark to determine performance of just using an empty var
// for zero value comparison
func BenchmarkValidString(b *testing.B) {
	uut := TestString{
		Required: "Required",
		Len:      "1234567890",
		Min:      "12345",
		Max:      "1231232",
	}
	var err error

	b.ResetTimer()
	for x := 0; x < b.N; x++ {
		err = uut.Validate()
	}
	errResult = err
}

// BenchmarkEmptyStruct is a quick benchmark to determine performance of just using an empty var
// for zero value comparison
func BenchmarkFailing1TestString(b *testing.B) {
	uut := TestString{
		Required: "",
		Len:      "1234567890",
		Min:      "12345",
		Max:      "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
	}
	var err error

	b.ResetTimer()
	for x := 0; x < b.N; x++ {
		err = uut.Validate()
	}
	errResult = err
}

// BenchmarkEmptyStruct is a quick benchmark to determine performance of just using an empty var
// for zero value comparison
func BenchmarkFailing2TestString(b *testing.B) {
	uut := TestString{
		Required: "",
		Len:      "12345678901",
		Min:      "12345",
		Max:      "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
	}
	var err error

	b.ResetTimer()
	for x := 0; x < b.N; x++ {
		err = uut.Validate()
	}
	errResult = err
}

// BenchmarkEmptyStruct is a quick benchmark to determine performance of just using an empty var
// for zero value comparison
func BenchmarkFailingAllTestString(b *testing.B) {
	uut := TestString{
		Required: "",
		Len:      "12345678901",
		Min:      "1234",
		Max:      "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901",
	}
	var err error

	b.ResetTimer()
	for x := 0; x < b.N; x++ {
		err = uut.Validate()
	}
	errResult = err
}

const (
	containsAnyPassingString = `Some string to do a contains any@ search# on!`
	containsAnyFailingString = `Some string to do a contains any search on.`
)

const (
	hexPassingString = `1234567890ABCDEF`
	hexFailingString = `1234567890ABCDEFG`
)

var benchmarks = []struct {
	name string
	uut  gencheck.Validateable
}{
	{"UUID Pass", TestUUID{UUID: uuid.NewV4().String()}},
	{"UUID Fail", TestUUID{UUID: uuid.NewV4().String() + "notauuid"}},
	{"Hex Pass", TestHex{Value: hexPassingString}},
	{"Hex Fail", TestHex{Value: hexFailingString}},
	{"ContainsAny Pass", TestContainsAny{Any: containsAnyPassingString}},
	{"ContainsAny Fail", TestContainsAny{Any: containsAnyFailingString}},
	{"TestStrings Pass", TestString{
		Required: "Required",
		Len:      "1234567890",
		Min:      "12345",
		Max:      "1231232",
	}},
	{"TestStrings Fail", TestString{
		Required: "",
		Len:      "12345678901",
		Min:      "1234",
		Max:      "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901",
	}},
	{"TestMap Pass", TestMap{Value: map[string]string{"test": "something"}}},
	{"TestMap Fail", TestMap{Value: map[string]string{"test2": "something"}}},
}

// BenchmarkCompareGencheckContainsAnyFail is a quick benchmark to determine performance of gencheck vs goplayground/validator
func BenchmarkCompareGencheck(b *testing.B) {
	var err error

	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err = bm.uut.Validate()
			}
			errResult = err
		})
	}
}

// BenchmarkCompareGencheckContainsAnyFail is a quick benchmark to determine performance of gencheck vs goplayground/validator
func BenchmarkComparePlayground(b *testing.B) {
	var err error
	validate := validator.New()

	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				err = validate.Struct(bm.uut)
			}
			errResult = err
		})
	}
}
