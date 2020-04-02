// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

//Kind ...
type Kind string

const (
	//NaT ...
	NaT = "NaT" // not a triangle
	//Equ ...
	Equ = "Equ" // equilateral
	//Iso ...
	Iso = "Iso" // isosceles
	//Sca ...
	Sca = "Sca" // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var k Kind

	if a+b >= c && b+c >= a && c+a >= b {
		if a == b && b == c {
			if a == 0 && b == 0 && c == 0 || (a < 0 || b < 0 || c < 0) {
				k = NaT
				return k
			}
			k = Equ
			return k
		} else if (a == b) || (b == c) || (c == a) {
			k = Iso
			return k
		} else if (a != b) || (b != c) || (c != a) {
			k = Sca
			return k
		}
	}
	k = NaT
	return k
}
