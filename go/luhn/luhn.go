package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

//Valid checks if the credit card number is valid or not
func Valid(number string) bool {
	//strip the spaces
	number = strings.Replace(number, " ", "", -1)
	l := len(number)

	//checks if number container characters apart from numbers
	r := regexp.MustCompile("^[0-9]*$")
	validNumber := r.Match([]byte(number))
	if !validNumber || l < 2 {
		return false
	}

	var toDouble bool
	var total int
	for i := l - 1; i >= 0; i-- {
		n, _ := strconv.Atoi(string(number[i]))
		if toDouble {
			double := n * 2
			if double > 9 {
				double = double - 9
			}
			total += double
			toDouble = false
		} else {
			total += n
			toDouble = true
		}
	}

	return total%10 == 0

}
