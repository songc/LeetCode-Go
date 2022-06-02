package main

//450. 删除二叉搜索树中的节点
//给定一个二叉搜索树的根节点 root 和一个值 key，删除二叉搜索树中的 key 对应的节点，并保证二叉搜索树的性质不变。返回二叉搜索树（有可能被更新）的根节点的引用。
//
//一般来说，删除节点可分为两个步骤：
//
//首先找到需要删除的节点；
//如果找到了，删除它。
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/delete-node-in-a-bst
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
func main() {

}

//Definition for a binary tree node.
//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}
	var pre *TreeNode
	curr := root
	for curr != nil {
		if curr.Val > key {
			pre = curr
			curr = curr.Left
		} else if curr.Val < key {
			pre = curr
			curr = curr.Right
		} else {
			break
		}
	}
	if curr == nil {
		return root
	}
	if curr.Left == nil && curr.Right == nil {
		curr = nil
	} else if curr.Left == nil {
		curr = curr.Right
	} else if curr.Right == nil {
		curr = curr.Left
	} else if curr.Right != nil && curr.Left != nil {
		soccer := curr.Right
		for soccer.Left != nil {
			soccer = soccer.Left
		}
		soccer.Left = curr.Left
		curr = curr.Right
	}
	if pre == nil {
		return curr
	}
	if pre.Left != nil && pre.Left.Val == key {
		pre.Left = curr
	} else {
		pre.Right = curr
	}
	return root
}
