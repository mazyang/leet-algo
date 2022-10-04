package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// 前序遍历
func (this *Codec) serialize(root *TreeNode) string {
	var builder strings.Builder
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			builder.WriteString("nil,")
			return
		}
		builder.WriteString(strconv.Itoa(node.Val) + ",")
		traverse(node.Left)
		traverse(node.Right)
	}
	traverse(root)
	return builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	vals := strings.Split(data, ",")
	var recur func() *TreeNode
	recur = func() *TreeNode {
		if len(vals) == 0 {
			return nil
		}
		first := vals[0]
		vals = vals[1:]
		if first == "nil" {
			return nil
		}
		val, _ := strconv.Atoi(first)
		// 确定根节点
		root := &TreeNode{Val: val}
		// 递归左右子树
		root.Left = recur()
		root.Right = recur()
		return root
	}
	return recur()
}

// 后序遍历
func (this *Codec) serialize2(root *TreeNode) string {
	var builder strings.Builder
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			builder.WriteString("nil,")
			return
		}
		traverse(node.Left)
		traverse(node.Right)
		builder.WriteString(strconv.Itoa(node.Val) + ",")
	}
	traverse(root)
	return builder.String()
}

func (this *Codec) deserialize2(data string) *TreeNode {
	vals := strings.Split(data[:len(data)-1], ",")
	var recur func() *TreeNode
	recur = func() *TreeNode {
		if len(vals) == 0 {
			return nil
		}
		last := vals[len(vals)-1]
		vals = vals[:len(vals)-1]
		if last == "nil" {
			return nil
		}
		val, _ := strconv.Atoi(last)
		// 确定根节点
		root := &TreeNode{Val: val}
		// 递归左右子树，注意递归顺序，先递归右子树，后递归左子树，因为是从后向前取列表元素
		root.Right = recur()
		root.Left = recur()
		return root
	}
	return recur()
}

func main() {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n1.Left = n2
	n1.Right = n3
	c := Codec{}
	a := c.serialize(n1)
	b := c.deserialize(a)
	println(b)
}
