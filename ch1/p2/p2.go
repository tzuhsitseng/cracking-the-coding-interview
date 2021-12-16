package main

import "fmt"

func main() {
	//first, second := "dog", "god"
	first, second := "dog", "dop"
	fmt.Println(checkPermutation(first, second))
}

func checkPermutation(first, second string) bool {
	// 如果兩個字串長度就不一樣，直接代表不是變位數
	if len(first) != len(second) {
		return false
	}

	// 將第一個字串所出現的字元作個計數
	checker := map[rune]int{}
	for _, c := range first {
		checker[c]++
	}

	// 繼續尋訪第二個字串
	// 若計數會變成 < 0 或是字元根本不存在的話，代表不是變位數
	for _, c := range second {
		cnt, ok := checker[c]
		if !ok || cnt == 0 {
			return false
		}
		if cnt == 1 {
			delete(checker, c)
		} else {
			checker[c]--
		}
	}

	// 若最後兩個字串的計數完全相抵消掉的話，那就是同位數了
	return len(checker) == 0
}
