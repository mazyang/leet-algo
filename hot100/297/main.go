package main

import (
	"fmt"
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

// 层次遍历
// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	builder := strings.Builder{}
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			builder.WriteString(strconv.Itoa(node.Val) + ",")
			// 无论是否为nil都加入队列中
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else {
			builder.WriteString("nil,")
		}
	}
	return builder.String()
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	vals := strings.Split(data[:len(data)-1], ",")
	rootVal, _ := strconv.Atoi(vals[0])
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) != 0 {
		// 每次从队列中弹出一个，i向前走两步，一个节点对应两个子节点
		node := queue[0]
		queue = queue[1:]
		if vals[i] != "nil" {
			val, _ := strconv.Atoi(vals[i])
			node.Left = &TreeNode{Val: val}
			queue = append(queue, node.Left)
		}
		i++
		if vals[i] != "nil" {
			val, _ := strconv.Atoi(vals[i])
			node.Right = &TreeNode{Val: val}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func main() {
	n1 := &TreeNode{Val: 1}
	n2 := &TreeNode{Val: 2}
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4}
	n5 := &TreeNode{Val: 5}
	n1.Left = n2
	n1.Right = n3
	n3.Left = n4
	n3.Right = n5
	ser := Constructor()
	//deser := Constructor()
	data := ser.serialize(n1)
	//ans := deser.deserialize(data);
	fmt.Println(data)

}
