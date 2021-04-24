package pascal

//Triangle calculates the pascals triangle
//It takes num as input which is number of rows for
//pascals triangle to calculate and returns the triangle
func Triangle(num int) [][]int {

	triangle := make([][]int, num)

	if num == 0 {
		return triangle
	}

	triangle[0] = []int{1}

	for row := 1; row < num; row++ {

		newRow := make([]int, row+1)
		prevRow := triangle[row-1]

		newRow[0] = 1

		for col := 1; col < row; col++ {
			newRow[col] = prevRow[col] + prevRow[col-1]
		}

		newRow[row] = 1
		triangle[row] = newRow
	}

	return triangle
}
