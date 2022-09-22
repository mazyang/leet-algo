package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 对称二叉树
func isSymmetric(root *TreeNode) bool {
	var recur func(*TreeNode, *TreeNode) bool
	recur = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		return recur(left.Left, right.Right) && recur(left.Right, right.Left)
	}
	return recur(root.Left, root.Right)
}

// 方法二：迭代
func isSymmetric2(root *TreeNode) bool {
	u, v := root.Left, root.Right
	var q []*TreeNode
	q = append(q, u)
	q = append(q, v)
	for len(q) > 0 {
		u, v = q[0], q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if u == nil || v == nil {
			return false
		}
		if u.Val != v.Val {
			return false
		}
		q = append(q, u.Left)
		q = append(q, v.Right)

		q = append(q, u.Right)
		q = append(q, v.Left)
	}
	return true
}
