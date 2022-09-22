package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var prev *TreeNode
	// first 和 second 是两个交换的节点，找到后交换
	var first *TreeNode
	var second *TreeNode

	var dfs func(root *TreeNode)
	// 中序遍历，prev保存上一个节点，与当前节点对比，判断是否小于当前节点，如果大于当前节点，说明当前节点是异常节点
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if prev != nil {
			if first == nil && prev.Val > root.Val {
				first = prev
			}
			if first != nil && prev.Val > root.Val {
				second = root
			}
		}
		prev = root
		dfs(root.Right)
	}
	dfs(root)
	first.Val, second.Val = second.Val, first.Val
}

func main() {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 3}
	n3 := &TreeNode{Val: 2}
	n2.Left = n1
	n2.Right = n3
	recoverTree(n2)
}
