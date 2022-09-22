package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 翻转二叉树
// 方法一：遍历, 将每个节点的左右子树交换
func invertTree(root *TreeNode) *TreeNode {
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		node.Left, node.Right = node.Right, node.Left
		traverse(node.Left)
		traverse(node.Right)
	}
	traverse(root)
	return root
}

// 方法二：分解, 分解为子树，每个子问题返回交换后的新子树
func invertTree2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := invertTree2(root.Left)
	right := invertTree2(root.Right)
	root.Left = right
	root.Right = left
	return root
}
