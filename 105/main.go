package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 从前序与中序遍历序列构造二叉树
// 方法一：递归
func buildTree(preorder []int, inorder []int) *TreeNode {
	hashTable := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		hashTable[inorder[i]] = i
	}
	var recur func(int, int, int, int) *TreeNode
	recur = func(preStart int, preEnd int, inStart int, inEnd int) *TreeNode {
		if preStart > preEnd {
			return nil
		}
		// 前序遍历的第一个元素就是根节点
		root := &TreeNode{Val: preorder[preStart]}
		// 如何快速获得根节点在中序遍历中的索引？：使用哈希表存放
		rootIdx := hashTable[root.Val]
		// rootIdx-inStart为左子树的长度  inEnd-rootIdx为右子树的长度
		root.Left = recur(preStart+1, preStart+1+rootIdx-inStart-1, inStart, rootIdx-1)
		root.Right = recur(preStart+1+rootIdx-inStart, preEnd, rootIdx+1, inEnd)
		return root
	}
	return recur(0, len(preorder)-1, 0, len(inorder)-1)
}

// 方法二：迭代
