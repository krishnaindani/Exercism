package isbn

import (
	"strconv"
	"strings"
)

//IsValidISBN ...
func IsValidISBN(input string) bool {
	s := strings.Split(input, "")
	// var wordTotal int
	var finalString []int
	for i := 0; i < len(s); i++ {
		if s[i] == "-" {
			continue
		} else if s[i] == "X" && i == len(s)-1 {
			finalString = append(finalString, 10)
		} else {
			value, err := strconv.Atoi(s[i])
			if err != nil {
				return false
			}
			finalString = append(finalString, value)
		}
	}

	if len(finalString) != 10 {
		return false
	}

	var total int
	for i := 0; i < len(finalString); i++ {
		total += (i + 1) * finalString[i]
	}

	return total%11 == 0
}
