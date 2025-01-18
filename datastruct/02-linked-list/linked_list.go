package linked_list

import "fmt"

type Node struct {
	Val  int32
	Next *Node
}

func NewNode(val int32, next *Node) *Node {
	node := &Node{Val: val, Next: next}
	return node
}

func (node *Node) InsertAfter(val int32) {
	node.Next = &Node{Val: val}
}

func (node *Node) InsertBefore(val int32) {
	before := &Node{Val: val, Next: node}
	node = before
}

func (node *Node) Delete(n *Node) {
	if n.Next != nil {
		node.Next = n.Next
		n = nil
	} else {
		node.Next = nil
	}
}

func (node *Node) ReverseByIteration() {
	temp := node
	for temp != nil {

	}

	node = node.Next

}

func (node *Node) Print() {
	fmt.Printf("%d -> ", node.Val)

	if node.Next != nil {
		node = node.Next
		node.Print()
	}
}
