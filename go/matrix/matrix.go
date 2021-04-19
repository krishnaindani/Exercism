package matrix

import (
	"fmt"
	"strings"
)

type Matrix [][]int

func New(input string) (*Matrix, error) {

	lines := strings.Split(input, "\n")
	for i, line := range lines {
		for _, e := range line {
			elements := strings.Split(string(e), " ")
			fmt.Printf("i: %v, elements: %v\n", i, elements)
		}

		fmt.Println("i", i)
		fmt.Println("line", line)
	}
	return &Matrix{}, nil
}

func (m *Matrix) Rows() [][]int {
	return nil
}

func (m *Matrix) Cols() [][]int {
	return nil
}

func (m *Matrix) Set(row, col, val int) bool {
	return false
}
