package benchmark

import "testing"

var result bool
var errResult error
var EmptyStructs = []TestString{
	TestString{},
	TestString{},
	TestString{},
	TestString{},
	TestString{},
	TestString{},
	TestString{},
	TestString{},
	TestString{},
	TestString{},
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
