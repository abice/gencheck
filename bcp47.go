package gencheck

import "golang.org/x/text/language"

// IsBCP47 will return an error if the string is not able to be parsed as a
// language descriptor.
// http://www.unicode.org/reports/tr35/#Unicode_Language_and_Locale_Identifiers
func IsBCP47(s *string) error {
	if s == nil {
		return nil
	}

	_, err := language.Parse(*s)

	return err
}
