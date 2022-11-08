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

// 反转链表的前N个节点
var successor *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		// 记录第 n + 1 个节点
		head.Next = successor
		// 返回新的头节点
		return head
	}
	newHead := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = successor
	return newHead
}

// 反转链表的一部分
func reverseBetween(head *ListNode, m, n int) *ListNode {
	if m == 1 {
		return reverseN(head, n)
	}
	head.Next = reverseBetween(head.Next, m-1, n-1)
	return head
}

// K个一组反转链表
// 方法一：递归法
func reverseKGroup(head *ListNode, k int) *ListNode {
	var reverse func(*ListNode, *ListNode) *ListNode
	// 反转start end之间的元素，不包括end
	reverse = func(start *ListNode, end *ListNode) *ListNode {
		var prev, cur, next *ListNode = nil, start, end
		for cur != end {
			next = cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}
		return prev
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		// 如果剩下的没有k个节点
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

// 方法二：迭代法
func reverseKGroup2(head *ListNode, k int) *ListNode {
	var reverse func(*ListNode, *ListNode) (*ListNode, *ListNode)
	// 反转start end之间的元素，不包括end，返回反转之后的头节点和尾节点
	reverse = func(start *ListNode, end *ListNode) (*ListNode, *ListNode) {
		var prev, cur, next *ListNode = nil, start, end
		for prev != end {
			next = cur.Next
			cur.Next = prev
			prev = cur
			cur = next
		}
		return end, start
	}
	dummyNode := &ListNode{Next: head}
	// pre保存k个一组的链表的上一个节点
	pre := dummyNode

	for head != nil {
		tail := pre
		// tail 遍历到第k个节点
		for i := 0; i < k; i++ {
			tail = tail.Next
			// 如果剩下的不足k个节点
			if tail == nil {
				return dummyNode.Next
			}
		}
		// 保存第k+1个节点
		nex := tail.Next
		// 反转从head到tail
		head, tail = reverse(head, tail)
		// 指向新的头节点
		pre.Next = head
		tail.Next = nex
		pre = tail
		head = tail.Next
	}
	return dummyNode.Next
}
