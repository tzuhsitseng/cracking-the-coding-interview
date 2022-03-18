package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(checkOneEditAway("pale", "ple"))
	fmt.Println(checkOneEditAway("pales", "pale"))
	fmt.Println(checkOneEditAway("pale", "bale"))
	fmt.Println(checkOneEditAway("pale", "bae"))
}

func checkOneEditAway(s1, s2 string) bool {
	if len(s1) == len(s2) {
		return checkOneEditReplace(s1, s2)
	} else if math.Abs(float64(len(s1)-len(s2))) == 1 {
		if len(s1)-len(s2) > 0 {
			s1, s2 = s2, s1
		}
		return checkOneEditInsertOrRemove(s1, s2)
	}
	return false
}

func checkOneEditInsertOrRemove(s1, s2 string) bool {
	foundDiff := false
	i, j := 0, 0

	for i < len(s1) && j < len(s2) {
		if s1[i] != s2[j] {
			if foundDiff {
				return false
			}
			foundDiff = true
			j++
		} else {
			i++
			j++
		}
	}

	return true
}

func checkOneEditReplace(s1, s2 string) bool {
	foundDiff := false

	for i, c := range s1 {
		if string(c) != string(s2[i]) {
			if foundDiff {
				return false
			}
			foundDiff = true
		}
	}

	return true
}
