package gencheck

import (
	"fmt"
	"regexp"
	"strings"
)

var hexMatcher = regexp.MustCompile("^[0-9a-fA-F]+$")

// IsHex validates that the given string is a hex value
func IsHex(s *string) error {
	if s == nil {
		return nil
	}

	matches := hexMatcher.MatchString(strings.TrimPrefix(*s, "0x"))
	if !matches {
		return fmt.Errorf("'%s' is not a hexadecimal string", *s)
	}

	return nil
}
