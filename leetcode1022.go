package main

//1022. 从根到叶的二进制数之和
//给出一棵二叉树，其上每个结点的值都是0或1。每一条从根到叶的路径都代表一个从最高有效位开始的二进制数。
//例如，如果路径为0 -> 1 -> 1 -> 0 -> 1，那么它表示二进制数01101，也就是13。
//对树上的每一片叶子，我们都要找出从根到该叶子的路径所表示的数字。
//返回这些数字之和。题目数据保证答案是一个 32 位 整数。
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, curr int) int {
	if node == nil {
		return 0
	}
	ncurr := (curr << 1) + node.Val
	if node.Left == nil && node.Right == nil {
		return ncurr
	}
	return dfs(node.Left, ncurr) + dfs(node.Right, ncurr)
}
