package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// 找到中间节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 如果链表长度为奇数，slow指向下一个
	//if fast != nil {
	//	slow = slow.Next
	//}
	// 将后面的一半链表反转
	var prev, cur, next *ListNode = nil, slow, slow
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	for prev != nil && head != nil {
		if prev.Val != head.Val {
			return false
		}
		prev = prev.Next
		head = head.Next
	}
	return true
}
