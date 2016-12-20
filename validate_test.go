package gencheck

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ValidateTestSuite struct {
	suite.Suite
}

func TestValidateTestSuite(t *testing.T) {
	suite.Run(t, new(ValidateTestSuite))
}

type CanValidate struct {
	Valid bool
}

func (c CanValidate) Validate() error {
	if c.Valid {
		return nil
	}
	return ValidationErrors{
		NewFieldError("CanValidate", "Valid", "", fmt.Errorf("Valid when true")),
		NewFieldError("CanValidate", "Valid", "valid", nil),
	}
}

type NotValidateable struct {
	Whatever string
}

// TestCanValidate_Happy
func (s *ValidateTestSuite) TestCanValidate_Happy() {
	uut := &CanValidate{Valid: true}
	s.Require().Nil(Validate(uut))
}

// TestCanValidate_ValidationError
func (s *ValidateTestSuite) TestCanValidate_ValidationError() {
	uut := CanValidate{Valid: false}
	err := Validate(uut)
	s.Require().NotNil(err)

	if ve, ok := err.(ValidationErrors); ok {
		s.Equal("validation: field validation failed for 'CanValidate.Valid': Valid when true\nvalidation: field validation failed for 'CanValidate.Valid': tag='valid'", ve.Error())
		s.Require().Len(ve, 2, "Should have 2 errors in the CanValidate Error")

		fe := ve[0]
		s.Equal("CanValidate", fe.Struct())
		s.Equal("Valid", fe.Field())
		s.Equal("", fe.Tag())
		s.Equal("Valid when true", fe.Message())
		s.Equal("validation: field validation failed for 'CanValidate.Valid': Valid when true", fe.Error())

		fe = ve[1]
		s.Equal("CanValidate", fe.Struct())
		s.Equal("Valid", fe.Field())
		s.Equal("valid", fe.Tag())
		s.Equal("", fe.Message())
		s.Equal("validation: field validation failed for 'CanValidate.Valid': tag='valid'", fe.Error())

	} else {
		s.FailNow("Error returned should have been validation errors")
	}
}

// TestNotValidateable tests to see if the string passed in is nil
func (s *ValidateTestSuite) TestNotValidateable() {
	uut := &NotValidateable{}
	s.Require().Nil(Validate(uut))
}
