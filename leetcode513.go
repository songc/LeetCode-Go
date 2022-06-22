package main

//  Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	ans := queue[0].Val
	for len(queue) > 0 {
		newQueue := make([]*TreeNode, 0)
		for _, v := range queue {
			if v.Left != nil {
				newQueue = append(newQueue, v.Left)
			}
			if v.Right != nil {
				newQueue = append(newQueue, v.Right)
			}
		}
		queue = newQueue
		if len(queue) > 0 {
			ans = queue[0].Val
		}
	}
	return ans
}
