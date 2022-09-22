package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 路径总和
// 方法一：递归
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if targetSum == 0 && root.Left == nil && root.Right == nil {
		return true
	}
	// 只要有一个为true，结果就为true
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

// 方法二：广度优先搜索
func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	queNode := []*TreeNode{root}
	// queVal 保存根节点到当前节点的路径和
	queVal := []int{root.Val}
	for len(queNode) != 0 {
		now := queNode[0]
		queNode = queNode[1:]
		temp := queVal[0]
		queVal = queVal[1:]
		if now.Left == nil && now.Right == nil {
			if temp == targetSum {
				return true
			}
			continue
		}
		if now.Left != nil {
			queNode = append(queNode, now.Left)
			queVal = append(queVal, now.Left.Val+temp)
		}
		if now.Right != nil {
			queNode = append(queNode, now.Right)
			queVal = append(queVal, now.Right.Val+temp)
		}
	}
	return false
}
