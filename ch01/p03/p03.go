package main

import "fmt"

func main() {
	emptySlice := make([]byte, 4)
	input := append([]byte("Mr John Smith"), emptySlice...)
	fmt.Println(string(URLify(input, 13)))
}

func URLify(input []byte, trueLength int) []byte {
	i, j := len(input)-1, trueLength-1

	for i > 0 && j > 0 && i >= j {
		if string(input[j]) == " " {
			input[i] = '0'
			input[i-1] = '2'
			input[i-2] = '%'
			i -= 3
		} else {
			input[i] = input[j]
			i--
		}
		j--
	}

	return input
}
