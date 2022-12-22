package main

type Node struct {
	Val  int
	Next *Node
}

func removeNode(head *Node, n int) *Node {
	dummy := &Node{Next: head}
	fast, slow := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

func main() {

}
