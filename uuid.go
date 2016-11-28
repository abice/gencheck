package gencheck

import (
	"fmt"
	"regexp"
	"strings"
)

var uuidMatcher = regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
var uuid3Matcher = regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$")
var uuid4Matcher = regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$")
var uuid5Matcher = regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$")

// IsUUID validates that the given string is a UUID value
func IsUUID(s *string) error {
	if s == nil {
		return nil
	}

	matches := uuidMatcher.MatchString(strings.ToLower(*s))
	if !matches {
		return fmt.Errorf("'%s' is not a UUID", *s)
	}

	return nil
}

// IsUUIDv3 validates that the given string is a UUIDv3 value
func IsUUIDv3(s *string) error {
	if s == nil {
		return nil
	}

	matches := uuid3Matcher.MatchString(strings.ToLower(*s))
	if !matches {
		return fmt.Errorf("'%s' is not a UUIDv3", *s)
	}

	return nil
}

// IsUUIDv4 validates that the given string is a UUIDv4 value
func IsUUIDv4(s *string) error {
	if s == nil {
		return nil
	}

	matches := uuid4Matcher.MatchString(strings.ToLower(*s))
	if !matches {
		return fmt.Errorf("'%s' is not a UUIDv4", *s)
	}

	return nil
}

// IsUUIDv5 validates that the given string is a UUIDv5 value
func IsUUIDv5(s *string) error {
	if s == nil {
		return nil
	}

	matches := uuid5Matcher.MatchString(strings.ToLower(*s))
	if !matches {
		return fmt.Errorf("'%s' is not a UUIDv5", *s)
	}

	return nil
}
