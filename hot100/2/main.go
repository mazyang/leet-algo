package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 虚拟头节点
	l3 := &ListNode{Val: -1}
	p := l3
	// 进位
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		val := carry
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		carry = val / 10
		val = val % 10
		p.Next = &ListNode{Val: val}
		p = p.Next
	}
	return l3.Next
}
