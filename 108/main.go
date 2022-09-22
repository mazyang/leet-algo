package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 将有序数组转为二叉搜索树
func sortedArrayToBST(nums []int) *TreeNode {
	var recur func(int, int) *TreeNode
	recur = func(start int, end int) *TreeNode {
		if start > end {
			return nil
		}
		mid := (start + end) / 2
		root := &TreeNode{Val: nums[mid]}
		root.Left = recur(start, mid-1)
		root.Right = recur(mid+1, end)
		return root
	}
	return recur(0, len(nums)-1)
}

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a[1:])
}
