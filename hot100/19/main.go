package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 双指针+虚拟头节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1}
	dummy.Next = head
	fast, slow := head, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	// 找到倒数第n+1个节点，fast=nil 的时候slow刚好在倒数第n+1个节点上
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
