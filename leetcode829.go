package main

func consecutiveNumbersSum(n int) int {
	isContruct := func(n int, k int) bool {
		if k%2 == 1 {
			return n%k == 0
		}
		return n%k > 0 && 2*n%k == 0
	}
	ans := 0
	for k := 1; k*(k+1) <= 2*n; k += 1 {
		if isContruct(n, k) {
			ans += 1
		}
	}
	return ans
}
