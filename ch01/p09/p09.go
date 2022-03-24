package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isRotation("waterbottle", "erbottlewat"))
	fmt.Println(isRotation("waterbottle", "rbottlewate"))
	fmt.Println(isRotation("waterbottle", "bottlewater"))
	fmt.Println(isRotation("waterbottle", "bottelwater"))
}

func isRotation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	if s1 == s2 {
		return true
	}

	// 假設 s2 為 s1 的 substring -> s1 = x + y, s2 = y + x
	// 兩個 s1 = x + y + x + y
	// 所以 s2 必定為兩個 s1 的 substring
	ss := s1 + s1
	return strings.Contains(ss, s2)
}
