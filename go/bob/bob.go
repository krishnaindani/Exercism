package bob

import (
	"strings"
	"unicode"
)

// Hey returns the response for the conversation
func Hey(remark string) string {

	remark = strings.TrimSpace(remark)

	switch {
	case len(remark) == 0:
		return "Fine. Be that way!"
	case strings.ToUpper(remark) == remark && remark[len(remark)-1] == '?' && containsCharacter(remark):
		return "Calm down, I know what I'm doing!"
	case strings.ToUpper(remark) == remark && containsCharacter(remark):
		return "Whoa, chill out!"
	case remark[len(remark)-1] == '?':
		return "Sure."
	default:
		return "Whatever."
	}
}

func containsCharacter(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) {
			return true
		}
	}
	return false
}
