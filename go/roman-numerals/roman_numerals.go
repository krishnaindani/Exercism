package romannumerals

import "fmt"

func ToRomanNumeral(number int) (string, error)  {

	if number <= 0 || number > 3000 {
		return "", fmt.Errorf("number %d not supported for roman conversion", number)
	}

	romanNum := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanSymbol := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var output string
	 i := 0

	for number > 0 {
		div := number/romanNum[i]
		number = number%romanNum[i]

		for div > 0 {
			output += romanSymbol[i]
			div--
		}
		i++
	}

	return output, nil
}