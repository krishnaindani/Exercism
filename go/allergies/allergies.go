package allergies

import "math"

var substances = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

//Allergies take score and returns what all you are allergic to
func Allergies(score uint) []string {
	var product []string

	for i, sub := range substances {
		if v := uint(math.Pow(2, float64(i))); v&score != 0 {
			product = append(product, sub)
		}
	}

	return product
}

//AllergicTo takes score and substance and returns if you are allergic to that
func AllergicTo(score uint, substance string) bool {

	index := -1
	for i, v := range substances {
		if v == substance {
			index = i
		}
	}

	if index == -1 {
		return false
	}

	v := uint(math.Pow(2, float64(index)))

	return v&score != 0
}
