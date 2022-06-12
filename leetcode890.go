package main

import "fmt"

func main() {
	words := []string{"abc", "deq", "mee", "aqq", "dkd", "ccc"}
	pattern := "abb"
	fmt.Print(findAndReplacePattern(words, pattern))

}

func findAndReplacePattern(words []string, pattern string) []string {
	target := make(map[string]bool)
	ans := make([]string, 0)
	for i := 0; i < len(pattern)-1; i++ {
		for j := i + 1; j < len(pattern); j++ {
			key := fmt.Sprintf("%d+%d", i, j)
			target[key] = (pattern[i] == pattern[j])
		}
	}
	for _, word := range words {
		flag := true
	re:
		for i := 0; i < len(word); i++ {
			for j := i + 1; j < len(word); j++ {
				key := fmt.Sprintf("%d+%d", i, j)
				if (word[i] == word[j]) != target[key] {
					flag = false
					break re
				}
			}
		}
		if flag {

			ans = append(ans, word)
		}
	}

	return ans
}
