package main

import (
	"strconv"
	"strings"
)

//449. 序列化和反序列化二叉搜索树
//序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。
//
//设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序列化为最初的二叉搜索树。
//
//编码的字符串应尽可能紧凑。
//
//
//
//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/serialize-and-deserialize-bst
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func main() {

}

//Definition for a binary tree node.
//type TreeNode struct {
//	Val int
//	Left *TreeNode
//	Right *TreeNode
//}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	stringBuild := make([]string, 0, 100)
	var dfsFunc func(node *TreeNode)
	dfsFunc = func(node *TreeNode) {
		if node == nil {
			return
		}
		stringBuild = append(stringBuild, strconv.Itoa(node.Val))
		dfsFunc(node.Left)
		dfsFunc(node.Right)
	}
	dfsFunc(root)
	return strings.Join(stringBuild, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	datalist := strings.Split(data, ",")
	numList := make([]int, 0, len(datalist))
	for _, s := range datalist {
		parseInt, err := strconv.Atoi(s)
		if err != nil {
			return nil
		}
		numList = append(numList, parseInt)
	}
	return construct(numList)
}

func construct(valueList []int) *TreeNode {
	if len(valueList) == 0 {
		return nil
	}
	ans := &TreeNode{Val: valueList[0]}
	i := 0
	for ; i < len(valueList); i++ {
		if valueList[i] > valueList[0] {
			break
		}
	}
	ans.Left = construct(valueList[1:i])
	if i != len(valueList) {
		ans.Right = construct(valueList[i:])
	}
	return ans
}
