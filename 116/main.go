package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	var traverse func(*Node, *Node)
	traverse = func(node1 *Node, node2 *Node) {
		if node1 == nil && node2 == nil {
			return
		}
		node1.Next = node2
		traverse(node1.Left, node1.Right)
		traverse(node2.Left, node2.Right)
		traverse(node1.Right, node2.Left)
	}
	traverse(root.Left, root.Right)
	return root
}
