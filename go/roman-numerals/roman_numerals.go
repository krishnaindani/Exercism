package romannumerals

//ToRomanNumeral converts deciman number to roman number
func ToRomanNumeral(number int) (string, error) {

	values := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4, 1,
	}

	symbols := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV",
		"I"}

	roman := ""
	i := 0

	for number > 0 {
		k := number / values[0]
		for j := 0; j < k; j++ {
			roman += symbols[i]
			number -= values[i]
		}
		i++
	}
	return roman, nil
}
