//go:generate gencheck -f=types.go

package benchmark

type TestString struct {
	Required string `valid:"required" validate:"required"`
	Len      string `valid:"len=10" validate:"len=10"`
	Min      string `valid:"min=5" validate:"min=5"`
	Max      string `valid:"max=100" validate:"max=100"`
}

type SingleString struct {
	Entry string `valid:"required"`
}

type TestUUID struct {
	UUID string `valid:"required,UUID" validate:"required,uuid"`
}

type TestContainsAny struct {
	Any string `valid:"containsany=@#!" validate:"containsany=@#!"`
}

type TestHex struct {
	Value string `valid:"hex" validate:"hexadecimal"`
}

type TestMap struct {
	Value map[string]string `valid:"contains=test" validate:"contains=test"`
}
