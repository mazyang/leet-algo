package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 层次遍历
// 方法一：广度优先搜索
func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var q = []*TreeNode{root}
	for len(q) > 0 {
		var t []int
		// 当队列的长度在不断变化的时候可以从后往前遍历
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

// 方法二：深度优先搜索
func levelOrder2(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		// 说明这是新的一层，需要创建数组
		if depth == len(ans) {
			ans = append(ans, make([]int, 0))
		}
		ans[depth] = append(ans[depth], root.Val)
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	return ans
}
