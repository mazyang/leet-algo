package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 判断两个二叉树是否相同
// 方法一：深度优先搜索
func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 相当于 p == nil && q == nil
	if p == q {
		return true
	}
	// 相当于 p==nil && q!=nil || p!=nil && q==nil
	if p == nil || q != nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 方法二：广度优先搜索
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == q {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	q1 := []*TreeNode{p}
	q2 := []*TreeNode{q}
	for len(q1) > 0 && len(q2) > 0 {
		// 从队列中取出第一个节点
		p, q = q1[0], q2[0]
		if p.Val != q.Val {
			return false
		}
		q1, q2 = q1[1:], q2[1:]
		la, ra := p.Left, p.Right
		lb, rb := q.Left, q.Right
		// 判断当前节点的左、右节点是否相同
		if (la != nil && lb == nil) || (lb != nil && la == nil) {
			return false
		}
		if (ra != nil && rb == nil) || (rb != nil && ra == nil) {
			return false
		}
		if la != nil {
			q1 = append(q1, la)
			q2 = append(q2, lb)
		}
		if ra != nil {
			q1 = append(q1, ra)
			q2 = append(q2, rb)
		}
	}
	return true
}
