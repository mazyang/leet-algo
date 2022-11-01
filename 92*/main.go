package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 方法一：递归法，left和right向前移动，直到left=1，反转链表的前right个节点
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

var successor *ListNode

// reverseN 将链表的前 n 个节点反转
func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		successor = head.Next
		return head
	}
	newHead := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return newHead
}

// 方法二：迭代法
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	// 设置 dummyNode 是这一类问题的一般做法
	dummyNode := &ListNode{Val: -1}
	dummyNode.Next = head
	pre := dummyNode
	for i := 0; i < left-1; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return dummyNode.Next
}
