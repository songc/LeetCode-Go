package main

func main() {

}

func minEatingSpeed(piles []int, h int) int {
	getNum := func(n int) int {
		ans := 0
		for _, pile := range piles {
			ans += (pile + n - 1) / n
		}
		return ans
	}
	left := 1
	right := 1
	for _, pile := range piles {
		if pile > right {
			right = pile
		}
	}
	for left < right {
		mid := (left + right) / 2
		target := getNum(mid)
		if target > h {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left

}
