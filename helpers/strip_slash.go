package helpers

import (
	"regexp"
)

// StripSlash removes slash (/) character from string if it is in a last position of the string.
func StripSlash(str string) string {
	re := regexp.MustCompile("/+$")
	return re.ReplaceAllString(str, "")
}
