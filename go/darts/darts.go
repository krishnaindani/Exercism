package darts

import "math"

//Score will calculate the score on where the dart hits
//outside target will give 0 point
//outside circle of target will give 1 point
//middle circle of target will give 5 points
//inner circle of target will give 10 points
func Score(x, y float64) int {

	center := []float64{0, 0}

	if inOrOnCircle(center, []float64{x, y}, 1) {
		return 10
	}

	if inOrOnCircle(center, []float64{x, y}, 5) {
		return 5
	}

	if inOrOnCircle(center, []float64{x, y}, 10) {
		return 1
	}

	return 0
}

//inOrOnCircle takes center and coordinates in x,y format with float64 type
//and checks if the coordinates are in or on circle
func inOrOnCircle(center, coordinates []float64, radius float64) bool {
	return (math.Pow(coordinates[0]-center[0], 2) +
		math.Pow(coordinates[1]-center[1], 2)) <=
		math.Pow(radius, 2)
}
