// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import "math"

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
	if (a <= 0 || b <= 0 || c <= 0) || (math.IsNaN(a) || math.IsInf(a, 1) || math.IsNaN(b) || math.IsInf(b, 1) || math.IsNaN(c) || math.IsInf(c, 1)) {
		return NaT
	}

	if a+b >= c && b+c >= a && c+a >= b {
		switch {
		case a == b && b == c && c == a:
			return Equ
		case (a == b) || (b == c) || (c == a):
			return Iso
		case (a != b) || (b != c) || (c != a):
			return Sca
		}
	}
	return NaT
}
