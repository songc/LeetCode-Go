package main

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	vmap := make(map[int]int)
	dfsTree(root, vmap)
	ans := make([]int, 0)
	target := 1
	for _, v := range vmap {
		if v > target {
			target = v
		}
	}
	for k, v := range vmap {
		if v == target {
			ans = append(ans, k)
		}
	}
	return ans
}

func dfsTree(node *TreeNode, vmap map[int]int) int {
	if node == nil {
		return 0
	}
	left := dfsTree(node.Left, vmap)
	right := dfsTree(node.Right, vmap)
	pre := left + right + node.Val
	if i, ok := vmap[pre]; ok {
		vmap[pre] = i + 1
	} else {
		vmap[pre] = 1
	}
	return pre
}
