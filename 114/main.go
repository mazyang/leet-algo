package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var traverse func(*TreeNode) *TreeNode
	traverse = func(node *TreeNode) *TreeNode {
		if node == nil {
			return nil
		}
		left := traverse(node.Left)
		right := traverse(node.Right)
		node.Left = nil
		node.Right = left
		tmp := node
		for tmp.Right != nil {
			tmp = tmp.Right
		}
		tmp.Right = right
		return node
	}
	traverse(root)
}
