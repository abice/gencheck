package benchmark

import (
	"strings"
	"testing"
	"time"

	"github.com/abice/gencheck"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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

var dummyString = "nothing but emptiness"

var benchmarks = []struct {
	name     string
	hasError bool
	uut      gencheck.Validateable
}{
	{"UUID Pass", false, TestUUID{UUID: uuid.New().String()}},
	{"UUID Fail", true, TestUUID{UUID: uuid.New().String() + "notauuid"}},
	{"Hex Pass", false, TestHex{Value: hexPassingString}},
	{"Hex Fail", true, TestHex{Value: hexFailingString}},
	{"ContainsAny Pass", false, TestContainsAny{Any: containsAnyPassingString}},
	{"ContainsAny Fail", true, TestContainsAny{Any: containsAnyFailingString}},
	{"TestStrings Pass", false, TestString{
		Required: "Required",
		Len:      "1234567890",
		Min:      "12345",
		Max:      "1231232",
	}},
	{"TestStrings Fail", true, TestString{
		Required: "",
		Len:      "12345678901",
		Min:      "1234",
		Max:      "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901",
	}},
	{"TestMap Pass", false, TestMap{Value: map[string]string{"test": "something"}}}, //Playground doesn't support this test.
	{"TestMap Fail", true, TestMap{Value: map[string]string{"test2": "something"}}},
	{"TestDive Pass", false, TestDive{Value: &SingleString{"This is a test"}}},
	{"TestDive Fail", true, TestDive{Value: &SingleString{""}}},
	{"TestDive Nil", true, TestDive{}},
	{"TestAll Pass", false, TestAll{
		Required: "Required",
		Len:      "1234567890",
		Min:      "12345",
		Max:      "1231232",
		CIDR:     "192.168.1.0/24",
		LteTime:  time.Now().Add(-1 * time.Second),
		GteTime:  time.Now().Add(60 * time.Second), // Make this a large number so that it's still after now when the test comes around.
		Gte:      1.23451,
		NotNil:   &dummyString,
		Contains: "The quick brown fox jumped over the lazy dog",
		Hex:      "1234567890AbCdEf",
		UUID:     uuid.New().String(),
		MinInt:   12345,
		MaxInt:   12345,
		URL:      "http://test.com/health?whatislife=something",
		URI:      "http://test.com/health?whatislife=something",
		Dive:     &SingleString{Entry: "This is a test"},
	}},
	{"TestAll Fail", true, TestAll{
		Required: "",
		Len:      "123456789",
		Min:      "1234",
		Max:      "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901",
		CIDR:     "192.168.1.0",
		LteTime:  time.Now().Add(60 * time.Second), // Make this a large number so that it's still after now when the test comes around.
		GteTime:  time.Now().Add(-1 * time.Second),
		Gte:      1.2344,
		NotNil:   nil,
		Contains: "The quick brown cat jumped over the lazy dog",
		Hex:      "1234567890AbCdEfG",
		UUID:     strings.ToUpper(uuid.New().String() + "adsf"),
		MinInt:   12344,
		MaxInt:   12346,
		URL:      "/health?whatislife=something",
		URI:      "?whatislife=something",
		Dive:     &SingleString{Entry: ""},
	}},
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
			if bm.hasError {
				assert.Error(b, err, "No Error when expected one.")
			} else {
				assert.NoError(b, err, "Error when there shouldn't be.")
			}
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
			if strings.HasPrefix(bm.name, "TestMap") {
				b.SkipNow()
			}
			for i := 0; i < b.N; i++ {
				err = validate.Struct(bm.uut)
			}
			errResult = err
			if bm.hasError {
				assert.Error(b, err, "No Error when expected one.")
			} else {
				assert.NoError(b, err, "Error when there shouldn't be.")
			}
		})
	}
}
