package main

import "fmt"

func main() {
	var matrix [][]int
	matrix = append(matrix, []int{0, 2, 3, 4})
	matrix = append(matrix, []int{5, 6, 7, 8})
	matrix = append(matrix, []int{9, 10, 0, 12})
	matrix = append(matrix, []int{13, 14, 15, 0})
	fmt.Println(setZeros(matrix))
}

func setZeros(matrix [][]int) [][]int {
	toZeroRows := map[int]bool{}
	toZeroCols := map[int]bool{}

	if len(matrix) == 0 {
		return matrix
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				toZeroRows[i] = true
				toZeroCols[j] = true
			}
		}
	}

	for k := range toZeroRows {
		setZerosByRow(matrix, k)
	}

	for k := range toZeroCols {
		setZerosByCol(matrix, k)
	}

	return matrix
}

func setZerosByRow(matrix [][]int, row int) [][]int {
	for i := range matrix[row] {
		matrix[row][i] = 0
	}
	return matrix
}

func setZerosByCol(matrix [][]int, col int) [][]int {
	for i := range matrix {
		matrix[i][col] = 0
	}
	return matrix
}
