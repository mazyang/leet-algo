package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 从中序和后序遍历序列构造二叉树
func buildTree(inorder []int, postorder []int) *TreeNode {
	dic := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		dic[inorder[i]] = i
	}
	var recur func(int, int, int, int) *TreeNode
	recur = func(inStart int, inEnd int, postStart int, postEnd int) *TreeNode {
		if postStart > postEnd {
			return nil
		}
		root := &TreeNode{Val: postorder[postEnd]}
		rootIdx := dic[root.Val]
		root.Left = recur(inStart, rootIdx-1, postStart, postStart+rootIdx-inStart-1)
		root.Right = recur(rootIdx+1, inEnd, postStart+rootIdx-inStart, postEnd-1)
		return root
	}
	return recur(0, len(inorder)-1, 0, len(postorder)-1)
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}
	buildTree(inorder, postorder)
}
