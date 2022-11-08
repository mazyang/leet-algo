package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 环形链表入环的第一个节点
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		// 代表slow和fast相遇
		if slow == fast {
			break
		}
	}
	// 说明没有环
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
