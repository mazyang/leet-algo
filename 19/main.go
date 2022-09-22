package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 删除链表的倒数第n个节点
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 虚拟头节点, 避免删除第一个节点
	res := &ListNode{Val: -1}
	res.Next = head
	fast := res
	// 需要找到倒数第n+1个节点
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	slow := res
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return res.Next
}
