//go:generate gencheck -f=types_test.go --failfast

package benchmark

type TestString struct {
	Required string `valid:"required"`
	Len      string `valid:"len=10"`
	Min      string `valid:"min=5"`
	Max      string `valid:"max=100"`
}

type SingleString struct {
	Entry string `valid:"required"`
}
