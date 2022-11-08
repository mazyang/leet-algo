package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 遇到p或q，直接返回
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// 遇到公共祖先
	if left != nil && right != nil {
		return root
	}
	// p q都不在以root为根的树中，返回nil
	if left == nil && right == nil {
		return nil
	}
	// p q中有一个在以root为根的树中，返回该节点
	if left == nil {
		return right
	} else {
		return left
	}
}
