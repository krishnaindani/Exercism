package luhn

import (
	"strconv"
	"strings"
)

//Valid checks if the credit card number is valid or not
func Valid(number string) bool {
	//strip the spaces
	number = strings.ReplaceAll(number, " ", "")
	l := len(number)
	if l < 2 {
		return false
	}

	var toDouble bool
	var total int
	for i := l - 1; i >= 0; i-- {
		n, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return false
		}
		if toDouble {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		total += n
		toDouble = !toDouble
	}
	return total%10 == 0

}
