package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var short string
	for _, v := range strings.FieldsFunc(s, func(c rune) bool {
		return c == '_' || c == ' ' || c == '-'
	}) {
		short += strings.ToUpper(string(v[0]))
	}
	return short
}
