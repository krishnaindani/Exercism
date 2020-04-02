package diffsquares

//SquareOfSum calculates the square of sum
func SquareOfSum(n int) (total int) {
	return (n * n * (n*n + 2*n + 1) / 4)
}

//SumOfSquares calculates the sum of squares
func SumOfSquares(n int) (total int) {
	return (n * (n + 1) * (2*n + 1) / 6)
}

//Difference calculates difference between SquareOfSum and SumOfSquares
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
