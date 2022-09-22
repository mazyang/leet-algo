package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 反转单链表
// 方法一：迭代法
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	var next *ListNode
	cur := head
	for cur != nil {
		next = cur.Next // 保存当前节点的下一个节点
		cur.Next = prev // 当前节点指向上一个节点
		prev = cur      // 上一个节点向前移动
		cur = next      // 当前节点向后移动
	}
	return prev
}

// 方法二：递归
func reverseList2(head *ListNode) *ListNode {
	if head == nil && head.Next == nil {
		return head
	}
	newHead := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
