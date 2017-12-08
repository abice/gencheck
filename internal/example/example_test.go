package example

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/abice/gencheck"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ExampleTestSuite: Run this test suite AFTER running gokay against example.go
type ExampleTestSuite struct {
	suite.Suite
}

// TestExampleTestSuite
func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

// TestValidateTestStruct_NoValues
func (s *ExampleTestSuite) TestValidateTestStruct_FailFast() {
	expected := gencheck.ValidationErrors{
		gencheck.NewFieldError("Test", "RequiredString", "required", fmt.Errorf("is required")),
	}

	underTest := Test{
		Embedded:    Embedded{FieldString: `1234`},
		MaxString:   "1234",
		MaxNumber:   1113.00001,
		MaxMultiple: []string{"", "", "", "", "", "", "", "", ""},
	}

	err := underTest.Validate()
	s.Require().IsType(gencheck.ValidationErrors{}, err, "Error returned was not ValidationErrors type")

	ve := err.(gencheck.ValidationErrors)

	s.Require().Equal(len(expected), len(ve), "Validation Errors were of different length")

	for i, fe := range ve {
		s.Equal(expected[i], fe)
	}
}

// TestValidateTestStruct_NoValues
func (s *ExampleTestSuite) TestValidateTestStruct_NoValues() {
	expected := gencheck.ValidationErrors{
		gencheck.NewFieldError("Test", "FieldString", "required", fmt.Errorf("is required")),
		gencheck.NewFieldError("Test", "RequiredMultiple", "required", fmt.Errorf("is required")),
		gencheck.NewFieldError("Test", "LenString", "len", fmt.Errorf("length mismatch")),
		gencheck.NewFieldError("Test", "LenNumber", "len", fmt.Errorf("length mismatch")),
		gencheck.NewFieldError("Test", "LenMultiple", "len", fmt.Errorf("length mismatch")),
		gencheck.NewFieldError("Test", "MinString", "min", fmt.Errorf("length failed check for min=1")),
		gencheck.NewFieldError("Test", "MinNumber", "min", fmt.Errorf("failed check for min=1113.00")),
		gencheck.NewFieldError("Test", "MinMultiple", "min", fmt.Errorf("length failed check for min=7")),
		gencheck.NewFieldError("Test", "MaxString", "max", fmt.Errorf("length failed check for max=3")),
		gencheck.NewFieldError("Test", "MaxNumber", "max", fmt.Errorf("failed check for max=1113.00")),
		gencheck.NewFieldError("Test", "MaxMultiple", "max", fmt.Errorf("length failed check for max=7")),
		gencheck.NewFieldError("Test", "EqString", "eq", fmt.Errorf("EqString did not equal 3")),
		gencheck.NewFieldError("Test", "EqNumber", "eq", fmt.Errorf("EqNumber did not equal 2.33")),
		gencheck.NewFieldError("Test", "EqMultiple", "eq", fmt.Errorf("Length of EqMultiple did not equal 7")),
		gencheck.NewFieldError("Test", "NeString", "ne", fmt.Errorf("NeString equaled ")),
		gencheck.NewFieldError("Test", "NeNumber", "ne", fmt.Errorf("NeNumber equaled 0")),
		gencheck.NewFieldError("Test", "NeMultiple", "ne", fmt.Errorf("Length of NeMultiple equaled")),
		gencheck.NewFieldError("Test", "LtString", "lt", fmt.Errorf("length failed check for lt=3")),
		gencheck.NewFieldError("Test", "LtNumber", "lt", fmt.Errorf("failed check for lt=5.56")),
		gencheck.NewFieldError("Test", "LtMultiple", "lt", fmt.Errorf("length failed check for lt=2")),
		gencheck.NewFieldError("Test", "LtTime", "lt", fmt.Errorf("is after now")),
		gencheck.NewFieldError("Test", "LteString", "lte", fmt.Errorf("length failed check for lte=3")),
		gencheck.NewFieldError("Test", "LteNumber", "lte", fmt.Errorf("failed check for lte=5.56")),
		gencheck.NewFieldError("Test", "LteMultiple", "lte", fmt.Errorf("length failed check for lte=2")),
		gencheck.NewFieldError("Test", "GtString", "gt", fmt.Errorf("length failed check for gt=3")),
		gencheck.NewFieldError("Test", "GtNumber", "gt", fmt.Errorf("failed check for gt=5.56")),
		gencheck.NewFieldError("Test", "GtMultiple", "gt", fmt.Errorf("length failed check for gt=2")),
		gencheck.NewFieldError("Test", "GtTime", "gt", fmt.Errorf("is before now")),
		gencheck.NewFieldError("Test", "GteString", "gte", fmt.Errorf("length failed check for gte=3")),
		gencheck.NewFieldError("Test", "GteNumber", "gte", fmt.Errorf("failed check for gte=5.56")),
		gencheck.NewFieldError("Test", "GteMultiple", "gte", fmt.Errorf("length failed check for gte=2")),
		gencheck.NewFieldError("Test", "GteTime", "gte", fmt.Errorf("is before now")),
		gencheck.NewFieldError("Test", "GteTimeVal", "gte", fmt.Errorf("is before %s", time.Now().UTC().Add(1*time.Second).Truncate(time.Second).Format("2006-01-02 15:04:05"))),
		gencheck.NewFieldError("Test", "GteTimePtr", "gte", fmt.Errorf("is before now")),
		gencheck.NewFieldError("Test", "HexadecimalString", "hexadecimal", fmt.Errorf("'' is not a hexadecimal string")),
		gencheck.NewFieldError("Test", "Contains", "contains", fmt.Errorf("Contains did not contain purpose")),
		gencheck.NewFieldError("Test", "ContainsPtr", "contains", fmt.Errorf("ContainsPtr did not contain purpose")),
		gencheck.NewFieldError("Test", "ContainsArray", "contains", fmt.Errorf("ContainsArray did not contain nonsense")),
		gencheck.NewFieldError("Test", "ContainsAny", "containsany", fmt.Errorf("ContainsAny did not contain any of !@#$")),
		gencheck.NewFieldError("Test", "UUID", "uuid", fmt.Errorf("'' is not a UUID")),
		gencheck.NewFieldError("Test", "UUID3", "uuid3", fmt.Errorf("'' is not a UUIDv3")),
		gencheck.NewFieldError("Test", "UUID4", "uuid4", fmt.Errorf("'' is not a UUIDv4")),
		gencheck.NewFieldError("Test", "UUID5", "uuid5", fmt.Errorf("'' is not a UUIDv5")),
		gencheck.NewFieldError("Test", "CIDR", "cidr", fmt.Errorf("invalid CIDR address")),
		gencheck.NewFieldError("Test", "CIDRv4", "cidrv4", fmt.Errorf("invalid CIDR address")),
		gencheck.NewFieldError("Test", "CIDRv6", "cidrv6", fmt.Errorf("invalid CIDR address")),
		gencheck.NewFieldError("Test", "MinIntPtr", "required", fmt.Errorf("is required")),
		gencheck.NewFieldError("Test", "InnerDive", "dive", fmt.Errorf("validation: field validation failed for 'Inner.EqCSFieldString': is required")),
		gencheck.NewFieldError("Test", "InnerDivePtr", "dive", fmt.Errorf("validation: field validation failed for 'Inner.EqCSFieldString': is required")),
		gencheck.NewFieldError("Test", "MapContains", "contains", fmt.Errorf("MapContains did not contain key")),
	}
	testTime := time.Now().UTC()
	notPurpose := "notPurpose"
	underTest := Test{
		RequiredString: "Here I am", // Put in required string to prevent fast failure
		MaxString:      "1234",
		MaxNumber:      1113.00001,
		MaxMultiple:    []string{"", "", "", "", "", "", "", "", ""},
		LtString:       "1234",
		LtNumber:       5.56000001,
		LtMultiple:     []string{"", "", ""},
		LtTime:         time.Now().UTC().Add(1 * time.Millisecond),
		LteString:      "1234",
		LteNumber:      5.56000001,
		LteMultiple:    []string{"", "", ""},
		GteTimePtr:     &testTime,
		ContainsPtr:    &notPurpose,
		InnerDivePtr:   &Inner{},
	}

	err := underTest.Validate()
	s.Error(err, "Error should not be nil")
	s.Require().IsType(gencheck.ValidationErrors{}, err, "Error returned was not ValidationErrors type")

	ve := err.(gencheck.ValidationErrors)

	s.Require().Equal(len(expected), len(ve), "Validation Errors were of different length")

	for i, fe := range ve {
		s.Contains(fe.Error(), strings.TrimSuffix(expected[i].Error(), "'"), "Error string mismatch")
		s.Equal(fe.Tag(), expected[i].Tag(), "Tag mismatch")
	}
}

// TestValidateTestStruct_Values
func (s *ExampleTestSuite) TestValidateTestStruct_Values() {
	i := int64(2000)
	underTest := Test{
		Embedded:          Embedded{FieldString: `1234`},
		LenMultiple:       []string{"", "", "", "", "", "", ""},
		LenNumber:         1113,
		LenString:         "a",
		EqMultiple:        []string{"", "", "", "", "", "", ""},
		EqNumber:          2.33,
		EqString:          "3",
		NeMultiple:        []string{"", "", "", "", "", "", ""},
		NeNumber:          2.33,
		NeString:          "3",
		RequiredMultiple:  []string{},
		RequiredString:    "b",
		MinString:         "1234567",
		MinNumber:         1113.000001,
		MinMultiple:       []string{"", "", "", "", "", "", "", ""},
		UUID:              "7112EE37-3219-4A26-BA01-1D230BC9257B",
		UUID3:             uuid.NewV3(uuid.NewV4(), "test").String(),
		UUID4:             strings.ToUpper(uuid.NewV4().String()),
		UUID5:             uuid.NewV5(uuid.NewV4(), "test").String(),
		MinIntPtr:         &i,
		GteString:         "1234",
		GteNumber:         5.5600001,
		GteMultiple:       []string{"", ""},
		LteTime:           time.Now().Add(-5 * time.Second),
		GtTime:            time.Now().Add(1 * time.Millisecond),
		GteTime:           time.Now().Add(5 * time.Second),
		GteTimeVal:        time.Now().Add(5 * time.Second),
		GtNumber:          5.5600001,
		GtMultiple:        []string{"", "", ""},
		GtString:          "1234",
		Contains:          "purpose Of this test",
		ContainsArray:     []string{"test", "last", "purpose", "nonsense"},
		ContainsAny:       "This is a test string!",
		InnerDive:         Inner{EqCSFieldString: "test"},
		InnerDivePtr:      &Inner{EqCSFieldString: "something"},
		MapContains:       map[string]interface{}{"key": "x"},
		HexadecimalString: "0x0F007BA11",
		CIDR:              "0.0.0.0/24",
		CIDRv4:            "0.0.0.0/24",
		CIDRv6:            "2620:0:2d0:200::7/32",
		URL:               "http://test.com",
		URI:               "scp://test.com/123",
	}

	err := underTest.Validate()
	s.NoError(err, "Valid Struct should not have had an error")

}

// TestValidateTestStruct_Values
func (s *ExampleTestSuite) TestValidateTestStruct_MinPtrFailure() {
	i := int64(1233)
	underTest := Test{
		Embedded:          Embedded{FieldString: `1234`},
		LenMultiple:       []string{"", "", "", "", "", "", ""},
		LenNumber:         1113,
		LenString:         "a",
		RequiredMultiple:  []string{},
		RequiredString:    "b",
		MinString:         "1234567",
		MinNumber:         1113.000001,
		MinMultiple:       []string{"", "", "", "", "", "", "", ""},
		UUID:              "7112EE37-3219-4A26-BA01-1D230BC9257B",
		UUID3:             uuid.NewV3(uuid.NewV4(), "test").String(),
		UUID4:             uuid.NewV4().String(),
		UUID5:             uuid.NewV5(uuid.NewV4(), "test").String(),
		GteString:         "1234",
		GteNumber:         5.5600001,
		GteMultiple:       []string{"", ""},
		LteTime:           time.Now().Add(-5 * time.Second),
		GtTime:            time.Now().Add(1 * time.Millisecond),
		GteTime:           time.Now().Add(5 * time.Second),
		GteTimeVal:        time.Now().Add(5 * time.Second),
		GtNumber:          5.5600001,
		GtMultiple:        []string{"", "", ""},
		GtString:          "1234",
		MinIntPtr:         &i,
		Contains:          "purpose Of this test",
		ContainsArray:     []string{"test", "last", "purpose", "nonsense"},
		ContainsAny:       "This is a test string!",
		InnerDive:         Inner{EqCSFieldString: "test"},
		InnerDivePtr:      &Inner{EqCSFieldString: "something"},
		MapContains:       map[string]interface{}{"key": "x"},
		HexadecimalString: "0x0F007BA11",
		CIDR:              "0.0.0.0/24",
		CIDRv4:            "0.0.0.0/24",
		CIDRv6:            "2620:0:2d0:200::7/32",
		EqMultiple:        []string{"", "", "", "", "", "", ""},
		EqNumber:          2.33,
		EqString:          "3",
		NeMultiple:        []string{"", "", "", "", "", "", ""},
		NeNumber:          2.33,
		NeString:          "3",
		URL:               "http://test.com",
		URI:               "scp://test.com",
	}

	err := underTest.Validate()
	s.Error(err, "Valid Struct should not have had an error")
	s.Require().IsType(gencheck.ValidationErrors{}, err, "Error returned was not ValidationErrors type")

	ve := err.(gencheck.ValidationErrors)
	s.Require().Len(ve, 1, "Should only have 1 validation error")

	s.Require().EqualValues(ve[0], gencheck.NewFieldError("Test", "MinIntPtr", "min", fmt.Errorf("failed check for min=1234")), "Error should be min error")
}

// TestValidateExample_NilMap
func (s *ExampleTestSuite) TestValidateExample_NilMap() {
	expected := gencheck.ValidationErrors{
		gencheck.NewFieldError("Example", "MapOfInterfaces", "notnil", fmt.Errorf("is Nil")),
	}

	underTest := Example{}

	err := underTest.Validate()
	s.EqualValues(expected, err)
}

// TestValidateNotNil_Map
func (s *ExampleTestSuite) TestValidateExample_Happy() {

	underTest := Example{
		MapOfInterfaces: make(map[string]interface{}),
	}

	err := underTest.Validate()
	s.Nil(err, "Valid Struct should not have had an error")
}

// TestValidateTestStruct_Values
func (s *ExampleTestSuite) TestValidateTestStruct_LteTime() {
	i := int64(1234)
	underTest := Test{
		Embedded:          Embedded{FieldString: `1234`},
		LenMultiple:       []string{"", "", "", "", "", "", ""},
		LenNumber:         1113,
		LenString:         "a",
		RequiredMultiple:  []string{},
		RequiredString:    "b",
		MinString:         "1234567",
		MinNumber:         1113.000001,
		MinMultiple:       []string{"", "", "", "", "", "", "", ""},
		UUID:              "7112EE37-3219-4A26-BA01-1D230BC9257B",
		UUID3:             uuid.NewV3(uuid.NewV4(), "test").String(),
		UUID4:             uuid.NewV4().String(),
		UUID5:             uuid.NewV5(uuid.NewV4(), "test").String(),
		GteString:         "1234",
		GteNumber:         5.5600001,
		GteMultiple:       []string{"", ""},
		LteTime:           time.Now().Add(1 * time.Second),
		GtTime:            time.Now().Add(1 * time.Millisecond),
		GteTime:           time.Now().Add(5 * time.Second),
		GteTimeVal:        time.Now().Add(5 * time.Second),
		GtNumber:          5.5600001,
		GtMultiple:        []string{"", "", ""},
		GtString:          "1234",
		MinIntPtr:         &i,
		Contains:          "purpose Of this test",
		ContainsArray:     []string{"test", "last", "purpose", "nonsense"},
		ContainsAny:       "This is a test string!",
		InnerDive:         Inner{EqCSFieldString: "test"},
		InnerDivePtr:      &Inner{EqCSFieldString: "something"},
		MapContains:       map[string]interface{}{"key": "x"},
		HexadecimalString: "0x0F007BA11",
		CIDR:              "0.0.0.0/24",
		CIDRv4:            "0.0.0.0/24",
		CIDRv6:            "2620:0:2d0:200::7/32",
		EqMultiple:        []string{"", "", "", "", "", "", ""},
		EqNumber:          2.33,
		EqString:          "3",
		NeMultiple:        []string{"", "", "", "", "", "", ""},
		NeNumber:          2.33,
		NeString:          "3",
		URL:               "http://test.com",
		URI:               "scp://test.com",
	}

	err := underTest.Validate()
	s.Error(err, "Valid Struct should have had an error")
	s.Require().IsType(gencheck.ValidationErrors{}, err, "Error returned was not ValidationErrors type")

	ve := err.(gencheck.ValidationErrors)
	s.Require().Len(ve, 1, "Should only have 1 validation error")

	s.Require().EqualValues(ve[0], gencheck.NewFieldError("Test", "LteTime", "lte", fmt.Errorf("is after now")), "Error should be lte time error")
}

// TestValidateTestStruct_Values
func TestValidateTestStruct_Individual(t *testing.T) {
	tests := map[string]struct {
		uut      Test
		field    string
		expected gencheck.FieldError
	}{
		"URL": {
			field:    "URL",
			expected: gencheck.NewFieldError("Test", "URL", "url", errors.New("parse x: invalid URI for request")),
			uut: Test{
				RequiredString: "x",
				URL:            "x",
			},
		},
		"URLNoScheme": {
			field:    "URL",
			expected: gencheck.NewFieldError("Test", "URL", "url", errors.New("URL is missing a scheme")),
			uut: Test{
				RequiredString: "x",
				URL:            "/x/123",
			},
		},
		"URI": {
			field:    "URI",
			expected: gencheck.NewFieldError("Test", "URI", "uri", errors.New("parse x: invalid URI for request")),
			uut: Test{
				RequiredString: "x",
				URI:            "x",
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(tt *testing.T) {
			err := test.uut.Validate()
			assert.Error(tt, err, "Valid Struct should have had an error")

			assert.IsType(tt, gencheck.ValidationErrors{}, err, "Error returned was not ValidationErrors type")

			ve := err.(gencheck.ValidationErrors)
			found := false

			for _, e := range ve {
				if e.Field() == test.field {
					found = true
					assert.EqualValues(tt, test.expected, e, "Error did not match expected")
				}
			}
			assert.True(tt, found, `Did not find expected error for field '%s'`, test.field)
		})
	}

}
