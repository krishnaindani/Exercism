package grains

import (
	"errors"
)

var errorInvalidIput = errors.New("invalid input, numbers should be between 1 to 64")

//Square ...
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errorInvalidIput
	}

	return 1 << (uint64(number - 1)), nil
}

//Total ...
func Total() uint64 {
	return (1 << 64) - 1
}
