package main

import "sort"

func main() {

}

func heightChecker(heights []int) int {
	target := make([]int, len(heights))
	copy(target, heights)
	sort.Ints(target)
	ans := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != target[i] {
			ans += 1
		}
	}
	return ans
}
