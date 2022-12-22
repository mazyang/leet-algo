package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	var sum int
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Right)
		sum += node.Val
		node.Val = sum
		traverse(node.Left)
	}
	traverse(root)
	return root
}
