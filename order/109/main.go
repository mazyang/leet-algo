package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 将有序链表转为二叉搜索树
// 方法一：双指针
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	fast := head
	slow := head
	pre := head
	// 找到中间节点, pre代表中间节点的前一个节点
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		pre = slow
		slow = slow.Next
	}
	pre.Next = nil
	// 当前slow节点就是根节点
	root := &TreeNode{Val: slow.Val}
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slow.Next)
	return root
}

// 方法二
func sortedListToBST2(head *ListNode) *TreeNode {
	var recur func(*ListNode, *ListNode) *TreeNode
	// 当前链表的左端点为left, 右端点为right, 左开右闭
	recur = func(left *ListNode, right *ListNode) *TreeNode {
		if left == right {
			return nil
		}
		fast, slow := left, left
		for fast != right && fast.Next != right {
			fast = fast.Next.Next
			slow = slow.Next
		}
		mid := slow
		root := &TreeNode{mid.Val, nil, nil}
		root.Left = recur(left, mid)
		root.Right = recur(mid.Next, right)
		return root
	}
	return recur(head, nil)
}

// 方法三：使用中序遍历
func sortedListToBST3(head *ListNode) *TreeNode {
	globalHead := head
	var buildTree func(int, int) *TreeNode
	buildTree = func(left int, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right + 1) / 2
		root := &TreeNode{}
		root.Left = buildTree(left, mid-1)
		root.Val = globalHead.Val
		globalHead = globalHead.Next
		root.Right = buildTree(mid+1, right)
		return root
	}
	length := 0
	for head != nil {
		length++
		head = head.Next
	}
	return buildTree(0, length-1)
}

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a[1:])
}
