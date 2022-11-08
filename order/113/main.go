package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 路径总和
// 方法一：回溯法
func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	var path []int
	var recur func(*TreeNode, int)
	recur = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}
		sum -= node.Val
		// 进入递归的时候将选择加入路径
		path = append(path, node.Val)
		// 撤销选择
		defer func() { path = path[:len(path)-1] }()
		if node.Left == nil && node.Right == nil && sum == 0 {
			res = append(res, append([]int{}, path...))
			// return 的时候也撤销选择
			return
		}
		recur(node.Left, sum)
		recur(node.Right, sum)
		// 也可以在最后撤销选择，那么上面的return就不需要了
		// path = path[:len(path)-1]
	}
	recur(root, targetSum)
	return res
}
