package main

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的锯齿形层序遍历
// 方法一：头插法
func zigzagLevelOrder(root *TreeNode) [][]int {
	var queue []*TreeNode
	var res [][]int
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) != 0 {
		var tmp []int
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			if len(res)%2 == 0 {
				tmp = append(tmp, node.Val)
			} else {
				// 在切片头部插入
				tmp = append(tmp, 0)
				copy(tmp[1:], tmp[0:])
				tmp[0] = node.Val
			}
			// 按照正常顺序入队列
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, tmp)
	}
	return res
}

// 方法二：每次循环打印两层
func zigzagLevelOrder2(root *TreeNode) [][]int {
	deque := list.New()
	var res [][]int
	if root != nil {
		deque.PushBack(root)
	}
	for deque.Len() != 0 {
		// 打印奇数层
		var tmp []int
		for i := deque.Len(); i > 0; i-- {
			// 从左向右打印
			node := deque.Front()
			deque.Remove(node)
			tmp = append(tmp, node.Value.(*TreeNode).Val)
			// 先左后右加入下层节点
			if node.Value.(*TreeNode).Left != nil {
				deque.PushBack(node.Value.(*TreeNode).Left)
			}
			if node.Value.(*TreeNode).Right != nil {
				deque.PushBack(node.Value.(*TreeNode).Right)
			}
		}
		res = append(res, tmp)
		// 若为空则提前跳出
		if deque.Len() == 0 {
			break
		}
		// 打印偶数层
		tmp = []int{}
		for i := deque.Len(); i > 0; i-- {
			// 从右向左打印
			node := deque.Back()
			deque.Remove(node)
			tmp = append(tmp, node.Value.(*TreeNode).Val)
			// 先右后左加入下层节点, 插入队列从队头插入, 得到的队列就是正常顺序
			if node.Value.(*TreeNode).Right != nil {
				deque.PushFront(node.Value.(*TreeNode).Right)
			}
			if node.Value.(*TreeNode).Left != nil {
				deque.PushFront(node.Value.(*TreeNode).Left)
			}
		}
		res = append(res, tmp)
	}
	return res
}
