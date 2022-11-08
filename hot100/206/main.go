package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转链表
// 方法一：迭代法
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur, next := head, head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// 方法二：递归
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
