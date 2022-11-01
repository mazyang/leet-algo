package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	q := []*TreeNode{root}
	for len(q) > 0 {
		var t []int
		// 当队列的长度在不断变化的时候可以从后往前遍历
		// 每次遍历一层, 然后记录在一个数组中
		for n := len(q); n > 0; n-- {
			node := q[0]
			// 出队列
			q = q[1:]
			t = append(t, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		ans = append(ans, t)
	}
	return ans
}
