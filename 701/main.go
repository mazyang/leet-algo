package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 找到空位置，插入新节点
	if root == nil {
		return &TreeNode{Val: val}
	}
	// 如果已经存在, 不需要重复插入，直接返回
	if root.Val == val {
		return root
	}
	// 在右子树或左子树寻找插入位置
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	} else {
		root.Left = insertIntoBST(root.Left, val)
	}
	return root
}
