package hamming

import "errors"

//
var (
	ErrUnequalLenths = errors.New("length is not equal")
)

//Distance calculates hamming distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, ErrUnequalLenths
	}
	if a == "" && b == "" {
		return 0, nil

	}

	if len(a) == len(b) {
		var differentWord int
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				differentWord++
			}
		}
		return differentWord, nil
	}
	return 0, nil
}
