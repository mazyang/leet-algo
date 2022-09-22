package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 自底向上的层序遍历
// 方法一：广度优先搜索
func levelOrderBottom(root *TreeNode) [][]int {
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
	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}
	return ans
}

// 方法二：深度优先搜索
func levelOrderBottom2(root *TreeNode) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		// 说明这是新的一层，需要创建数组插入ans的头部
		if depth == len(ans) {
			ans = append(ans, make([]int, 0))
			copy(ans[1:], ans[0:])
			ans[0] = make([]int, 0)
		}
		// 4, 3, 2, 1, 0 当前depth=1, len(ans)=5, 5-1-1=3
		ans[len(ans)-depth-1] = append(ans[len(ans)-depth-1], root.Val)
		dfs(root.Left, depth+1)
		dfs(root.Right, depth+1)
	}
	dfs(root, 0)
	return ans
}
