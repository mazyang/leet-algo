package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{Val: -1}
	p1, p2, p3 := list1, list2, res
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			p3.Next = p1
			p1 = p1.Next
		} else {
			p3.Next = p2
			p2 = p2.Next
		}
		p3 = p3.Next
	}
	if p1 == nil {
		p3.Next = p2
	}
	if p2 == nil {
		p3.Next = p1
	}
	return res.Next
}
