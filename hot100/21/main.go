package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 合并两个有序链表
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p1, p2 := list1, list2
	// 虚拟头节点
	dummy := &ListNode{Val: -1}
	p3 := dummy
	for p1 != nil || p2 != nil {
		if p1 == nil {
			p3.Next = p2
			break
		}
		if p2 == nil {
			p3.Next = p1
			break
		}
		if p1.Val < p2.Val {
			p3.Next = p1
			p3 = p3.Next
			p1 = p1.Next
		} else {
			p3.Next = p2
			p3 = p3.Next
			p2 = p2.Next
		}
	}
	return dummy.Next
}
