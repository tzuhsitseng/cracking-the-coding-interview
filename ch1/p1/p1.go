package main

import "fmt"

func main() {
	// 假設 input 會是 ascii code 的 scope
	input := "abc"
	//fmt.Println(checkUniqueCharset(input))
	fmt.Println(checkUniqueCharsetWithoutHashMap(input))
}

func checkUniqueCharset(input string) bool {
	// 若字串長度大於 ascii code 最大字元數 (128) 就一定是有重複
	if len(input) > 128 {
		return false
	}

	// 用個 map 來紀錄字元是否有出現過
	checker := map[rune]bool{}
	for _, ch := range input {
		if _, ok := checker[ch]; ok {
			return false
		}
		checker[ch] = true
	}

	return true
}

func checkUniqueCharsetWithoutHashMap(input string) bool {
	// 若字串長度大於 ascii code 最大字元數 (128) 就一定是有重複
	if len(input) > 128 {
		return false
	}

	// 用比較暴力的雙層迴圈來找出有沒有一樣的
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i] == input[j] {
				return false
			}
		}
	}

	return true
}
