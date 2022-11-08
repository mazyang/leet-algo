package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 获得所有的二叉搜索树
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	// allTrees 保存所有从 start 到 end 的二叉搜索树
	var allTrees []*TreeNode
	// 枚举可行根节点
	for i := start; i <= end; i++ {
		// 因为二叉搜索树左子树所有节点比根节点小，右子树比根节点大，因此构建 leftTrees 和 rightTrees
		// 获得所有可行的左子树集合
		leftTrees := helper(start, i-1)
		// 获得所有可行的右子树集合
		rightTrees := helper(i+1, end)
		// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
		// 固定左节点，遍历右节点
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currTree := &TreeNode{i, nil, nil}
				currTree.Left = left
				currTree.Right = right
				allTrees = append(allTrees, currTree)
			}
		}
	}
	return allTrees
}

func main() {
	fmt.Println(len(generateTrees(7)))
}
