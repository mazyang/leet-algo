package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 如果找到了p或q
	if root == p || root == q {
		return root
	}
	// 在左右两侧寻找 p q
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	// p q 在左右两侧找到
	if left != nil && right != nil {
		return root
	}
	// 一个都没有找到
	if left == nil && right == nil {
		return nil
	}
	// 找到一个
	if left != nil {
		return left
	} else {
		return right
	}
}
