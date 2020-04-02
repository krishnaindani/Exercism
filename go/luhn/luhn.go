package luhn

import (
	"fmt"
	"strings"
)

//Valid checks if the credit card number is valid or not
func Valid(number string) bool {
	number = strings.Replace(number, " ", "", -1)
	l := len(number)

	if l < 2 {
		return false
	}

	validNumber := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

	fmt.Println(validNumber)

	for _, char := range number {
		for _, v := range validNumber {
			if string(char) == v {
				continue
			}
		}
	}

	return false

}
