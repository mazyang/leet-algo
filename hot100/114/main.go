package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// 将左节点展平，此时root.Left 已经是一个链表了
	flatten(root.Left)
	flatten(root.Right)
	left := root.Left
	right := root.Right
	root.Left = nil
	root.Right = left
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}
