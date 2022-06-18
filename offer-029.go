package main

func main() {

}

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	node := &Node{x, nil}
	if aNode == nil {
		node.Next = node
		return node
	}
	if aNode.Next == aNode {
		node.Next = aNode
		aNode.Next = node
		return aNode
	}
	tmp := aNode
	for tmp.Next != aNode {
		if tmp.Val > tmp.Next.Val && (tmp.Next.Val > x || tmp.Val < x) {
			break
		}
		if tmp.Val <= x && tmp.Next.Val >= x {
			break
		}
		tmp = tmp.Next
	}
	node.Next = tmp.Next
	tmp.Next = node
	return aNode
}
