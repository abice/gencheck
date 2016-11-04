package gencheck

import (
	"fmt"
	"regexp"
	"strings"
)

var uuidMatcher = regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")

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
