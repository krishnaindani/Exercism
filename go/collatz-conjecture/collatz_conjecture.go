package collatzconjecture

import "errors"

var (
	errInvalidInput  = errors.New("Invalid Input")
	errNegativeInput = errors.New("Negative Input provided")
)

//CollatzConjecture ...
func CollatzConjecture(number int) (int, error) {
	var count int
	if number == 0 {
		return 0, errInvalidInput
	} else if number < 0 {
		return 0, errNegativeInput
	}

	for i := 0; number != 1; i++ {
		if number%2 == 0 {
			number /= 2
			count++
		} else {
			number = number*3 + 1
			count++
		}
	}
	return count, nil

}
