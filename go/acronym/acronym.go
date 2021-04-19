package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var short string

	f := func(c rune) bool {
		return c == '_' || c == ' ' || c == '-'
	}

	words := strings.FieldsFunc(s, f)

	for _, v := range words {
		short += strings.ToUpper(string(v[0]))
	}

	return short
}
