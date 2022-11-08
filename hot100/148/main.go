package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 排序链表
// 方法一：自顶向下 归并排序 空间复杂度O(n)
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(mid)
	res := &ListNode{Val: -1}
	tmp := res
	for left != nil && right != nil {
		if left.Val > right.Val {
			tmp.Next = right
			right = right.Next
		} else {
			tmp.Next = left
			left = left.Next
		}
		tmp = tmp.Next
	}
	if left != nil {
		tmp.Next = left
	}
	if right != nil {
		tmp.Next = right
	}
	return res.Next
}

// 方法二：自底向上 空间复杂度O(1)
func sortList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 1.使用length统计链表长度
	length := 0
	for node := head; node != nil; node = node.Next {
		length++
	}
	// 2.初始化，引入虚拟头节点
	dummyHead := &ListNode{Next: head}
	// subLength 表示每次需要排序的子链表的长度，初始为1，每次进行位运算，1 2 4 8... 合并子链表
	for subLength := 1; subLength < length; subLength <<= 1 {
		prev, cur := dummyHead, dummyHead.Next
		// 每次将链表拆分成若干个长度为 subLength 的子链表，然后两两合并
		for cur != nil {
			// head1代表第一个子链表
			head1 := cur
			for i := 1; i < subLength && cur.Next != nil; i++ {
				cur = cur.Next
			}
			// head2代表第二个子链表
			head2 := cur.Next
			cur.Next = nil // 断开连接
			cur = head2
			for i := 1; i < subLength && cur != nil && cur.Next != nil; i++ {
				cur = cur.Next
			}
			// 再次断开 第二个子链表最后的next的链接
			var next *ListNode
			if cur != nil {
				next = cur.Next
				cur.Next = nil
			}

			prev.Next = merge(head1, head2)
			// 将prev移动到合并之后的两个子链表后面
			for prev.Next != nil {
				prev = prev.Next
			}
			// 继续拆分子链表
			cur = next
		}
	}
	return dummyHead.Next
}

// 将两个有序链表合并
func merge(head1, head2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp, temp1, temp2 := dummyHead, head1, head2
	for temp1 != nil && temp2 != nil {
		if temp1.Val <= temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else if temp2 != nil {
		temp.Next = temp2
	}
	return dummyHead.Next
}
