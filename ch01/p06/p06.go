package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compressString("aabcccccaaa"))
	fmt.Println(compressString("abc"))
}

func compressString(s string) string {
	result := ""
	strLength := len(s)

	if strLength < 3 {
		return s
	}

	i := 0
	j := 0

	for i < strLength && j < strLength {
		j++

		if j == strLength || s[j] != s[i] {
			consecutiveCnt := j - i
			result += string(s[i]) + strconv.Itoa(consecutiveCnt)
			i = j
		}
	}

	if strLength <= len(result) {
		return s
	}

	return result
}
