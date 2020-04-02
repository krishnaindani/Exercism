package isogram

import "strings"

//IsIsogram checks whether a word is isogram or not and return bool.
func IsIsogram(s string) bool {
	var seen = make(map[rune]bool)

	for _, r := range strings.ToLower(s) {
		if r == ' ' || r == '-' {
			continue
		}
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
