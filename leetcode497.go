package main

import (
	"math/rand"
	"sort"
)

func main() {

}

type Solution struct {
	rects  [][]int
	preSum []int
}

func Constructor(rects [][]int) Solution {
	preSum := make([]int, 0, len(rects)+1)
	preSum = append(preSum, 0)
	pre := 0
	for _, rect := range rects {
		pre += (rect[2] - rect[0] + 1) * (rect[3] - rect[1] + 1)
		preSum = append(preSum, pre)
	}
	return Solution{rects: rects, preSum: preSum}
}

func (this *Solution) Pick() []int {
	num := rand.Intn(this.preSum[len(this.preSum)-1])
	ind := sort.SearchInts(this.preSum, num+1) - 1
	r := this.rects[ind]
	dx := rand.Intn(r[2] - r[0] + 1)
	dy := rand.Intn(r[3] - r[1] + 1)
	return []int{r[0] + dx, r[1] + dy}
}
