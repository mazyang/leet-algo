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

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	builder := strings.Builder{}
	builder.WriteString("[")
	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			builder.WriteString(strconv.Itoa(node.Val) + ",")
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		} else {
			builder.WriteString("nil,")
		}
	}
	res := builder.String()
	res = res[:len(res)-1]
	return res + "]"
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "[]" {
		return nil
	}
	vals := strings.Split(data[1:len(data)-1], ",")
	rootVal, _ := strconv.Atoi(vals[0])
	root := &TreeNode{Val: rootVal}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) != 0 {
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
