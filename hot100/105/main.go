package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 从前序和中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	hashMap := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		hashMap[inorder[i]] = i
	}
	var recur func(int, int, int, int) *TreeNode
	recur = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
		if preStart > preEnd {
			return nil
		}
		// 前序遍历第一个就是根节点
		root := &TreeNode{Val: preorder[preStart]}
		rootIdx := hashMap[preorder[preStart]]
		root.Left = recur(preStart+1, preStart+rootIdx-inStart, inStart, rootIdx)
		root.Right = recur(preStart+rootIdx-inStart+1, preEnd, rootIdx+1, inEnd)
		return root
	}
	return recur(0, len(preorder)-1, 0, len(inorder))
}
