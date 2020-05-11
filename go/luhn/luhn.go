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
	if l < 2 {
		return false
	}

	//checks if number container characters apart from numbers
	r := regexp.MustCompile("^[0-9]*$")
	validNumber := r.Match([]byte(number))
	if !validNumber {
		return false
	}

	//convert string to int
	numberInt := make([]int, l)
	for i := range number {
		numberInt[i], _ = strconv.Atoi(string(number[i]))
	}

	for i := l - 2; i >= 0; i = i - 2 {
		n := numberInt[i]
		double := n * 2
		if double > 9 {
			double = double - 9
		}
		numberInt[i] = double
	}
	//check is number is divisible
	flag := numberIsDivisible(numberInt)
	if !flag {
		return false
	}
	return true

}

func numberIsDivisible(n []int) bool {
	var total int
	for _, v := range n {
		total += v
	}
	if total%10 == 0 {
		return true
	}
	return false
}
