package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isPalindromePermutation("tact coa"))
	//fmt.Println(isPalindromePermutation("abc"))
}

func isPalindromePermutation(s string) bool {
	charsCnt := map[rune]int{}
	oddCnt := 0

	// remove all non-letter characters
	s = strings.ReplaceAll(s, " ", "")

	// count all characters
	for _, c := range s {
		charsCnt[c]++
	}

	// check no more than one character that is odd
	for _, cnt := range charsCnt {
		if cnt%2 != 0 {
			oddCnt++
		}
		if oddCnt > 1 {
			return false
		}
	}

	return true
}
