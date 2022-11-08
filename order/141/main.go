package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 环形链表
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		// 代表slow和fast相遇
		if slow == fast {
			return true
		}
	}
	return false
}
