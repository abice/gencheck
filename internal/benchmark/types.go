//go:generate gencheck -f=types.go

package benchmark

import "time"

type TestString struct {
	Required string `valid:"required" validate:"required"`
	Len      string `valid:"len=10" validate:"len=10"`
	Min      string `valid:"min=5" validate:"min=5"`
	Max      string `valid:"max=100" validate:"max=100"`
}

type SingleString struct {
	Entry string `valid:"required" validate:"required"`
}

type TestUUID struct {
	UUID string `valid:"required,uuid" validate:"required,uuid"`
}

type TestContainsAny struct {
	Any string `valid:"containsany=@#!" validate:"containsany=@#!"`
}

type TestHex struct {
	Value string `valid:"hex" validate:"hexadecimal"`
}

type TestMap struct {
	Value map[string]string `valid:"contains=test" validate:"contains=test"` //Playground doesn't really support this usecase.
}

type TestDive struct {
	Value *SingleString `valid:"required,dive" validate:"dive"` // Added required flag to have the same logic as playground
}

type TestAll struct {
	Required string        `valid:"required" validate:"required"`
	Len      string        `valid:"len=10" validate:"len=10"`
	Min      string        `valid:"min=5" validate:"min=5"`
	Max      string        `valid:"max=100" validate:"max=100"`
	CIDR     string        `valid:"required,cidr" validate:"required,cidr"`
	LteTime  time.Time     `valid:"lte" validate:"lte"`
	GteTime  time.Time     `valid:"gte" validate:"gte"`
	Gte      float64       `valid:"gte=1.2345" validate:"gte=1.2345"`
	NotNil   *string       `valid:"required" validate:"required"`
	Contains string        `valid:"contains=fox" validate:"contains=fox"`
	Hex      string        `valid:"hex" validate:"hexadecimal"`
	UUID     string        `valid:"uuid" validate:"uuid"`
	Dive     *SingleString `valid:"required,dive" validate:"dive"` // Added required flag to have the same logic as playground
}
