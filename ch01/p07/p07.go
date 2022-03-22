package main

import "fmt"

func main() {
	var data [][]int
	data = append(data, []int{1, 2, 3, 4})
	data = append(data, []int{5, 6, 7, 8})
	data = append(data, []int{9, 10, 11, 12})
	data = append(data, []int{13, 14, 15, 16})
	fmt.Println(rotate(data))
}

func rotate(matrix [][]int) [][]int {
	matrixLen := len(matrix)

	if matrixLen == 0 {
		return matrix
	}

	if matrixLen != len(matrix[0]) {
		return matrix
	}

	for l := 0; l < matrixLen/2; l++ {
		first := l
		last := matrixLen - 1 - l

		for i := first; i < last; i++ {
			offset := i - first

			// store top val as temp
			tmp := matrix[first][i]

			// store left val to top
			matrix[first][i] = matrix[last-offset][first]

			// store bottom val to left
			matrix[last-offset][first] = matrix[last][last-offset]

			// store right val to bottom
			matrix[last][last-offset] = matrix[i][last]

			// store tmp val (top val) to right
			matrix[i][last] = tmp
		}
	}

	return matrix
}
