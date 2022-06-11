package main

import "fmt"

func main() {
	s := "bccb"
	fmt.Println(countPalindromicSubsequences(s))
}

func countPalindromicSubsequences(s string) int {
	var mod int = 1e9 + 7
	dp := [4][][]int{}
	n := len(s)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j, _ := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}
	for i, c := range s {
		dp[c-'a'][i][i] = 1
	}
	tmps := "abcd"
	for sz := 2; sz <= n; sz++ {
		for i, j := 0, sz-1; j < n; i, j = i+1, j+1 {
			for k, _ := range tmps {
				c := tmps[k]
				if s[i] == c && s[j] == c {
					dp[c-'a'][i][j] = (2 + dp[0][i+1][j-1] + dp[1][i+1][j-1] + dp[2][i+1][j-1] + dp[3][i+1][j-1]) % mod
				} else if s[i] == c {
					dp[c-'a'][i][j] = dp[c-'a'][i][j-1]
				} else if s[j] == c {
					dp[c-'a'][i][j] = dp[c-'a'][i+1][j]
				} else {
					dp[c-'a'][i][j] = dp[c-'a'][i+1][j-1]
				}
			}
		}
	}
	ans := 0
	for _, d := range dp {
		ans += d[0][n-1]
	}
	return ans % mod
}
