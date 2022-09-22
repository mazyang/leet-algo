package main

// 双向链表

type Node struct {
	Val        int
	Prev, Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func Constructor() *LinkedList {
	head := &Node{Val: 0}
	tail := &Node{Val: 0}
	head.Next = tail
	tail.Prev = head
	return &LinkedList{
		Head: head,
		Tail: tail,
		Size: 0,
	}
}

func (this *LinkedList) PushFront(x *Node) {
	x.Next = this.Head.Next
	x.Prev = this.Head
	this.Head.Next.Prev = x
	this.Head.Next = x
	this.Size++
}

func (this *LinkedList) PushBack(x *Node) {
	x.Next = this.Tail
	x.Prev = this.Tail.Prev
	this.Tail.Prev.Next = x
	this.Tail.Prev = x
	this.Size++
}

func (this *LinkedList) Remove(x *Node) {
	x.Prev.Next = x.Next
	x.Next.Prev = x.Prev
	x.Prev = nil
	x.Next = nil
	this.Size--
}
