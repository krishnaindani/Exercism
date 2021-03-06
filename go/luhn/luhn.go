package luhn

import (
	"strconv"
	"strings"
)

//Valid checks if the credit card number is valid or not
func Valid(number string) bool {
	//strip the spaces
	number = strings.ReplaceAll(number, " ", "")
	if len(number) < 2 {
		return false
	}

	var toDouble bool
	var total int
	toDouble = len(number)%2 == 0

	for _, r := range number {
		n, err := strconv.Atoi(string(r))
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
