package main

import (
	"fmt"
)

//473. 火柴拼正方形
//你将得到一个整数数组 matchsticks ，其中 matchsticks[i] 是第 i 个火柴棒的长度。你要用 所有的火柴棍 拼成一个正方形。你 不能折断 任何一根火柴棒，但你可以把它们连在一起，而且每根火柴棒必须 使用一次 。
//
//如果你能使这个正方形，则返回 true ，否则返回 false 。
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/matchsticks-to-square
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
//
//1 <= matchsticks.length <= 15
//1 <= matchsticks[i] <= 108

func main() {
	matchsticks := []int{1, 1, 2, 2, 2}
	fmt.Print(makesquare(matchsticks))

}

func makesquare(matchsticks []int) bool {
	allLen := 0
	for _, matchstick := range matchsticks {
		allLen += matchstick
	}
	if allLen%4 != 0 {
		return false
	}
	target := allLen / 4
	dp := make([]int, 1<<len(matchsticks))
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
	}
	for s := 1; s < len(dp); s++ {
		for k, v := range matchsticks {
			if s>>k&1 == 0 {
				continue
			}
			s1 := s &^ (1 << k)
			if dp[s1] >= 0 && dp[s1]+v <= target {
				dp[s] = (dp[s1] + v) % target
				break
			}
		}
	}
	return dp[len(dp)-1] == 0
}
