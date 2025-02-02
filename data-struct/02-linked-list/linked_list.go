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

/*
*
迭代法：
使用三个指针：prev、current和next。
遍历链表，每次将current的Next指针指向prev，然后更新prev和current。
最后返回prev，即新的头节点
*/
func (node *Node) ReverseByIteration() *Node {
	var prev *Node
	current := node
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}

/*
*
递归法：
递归地翻转当前节点之后的链表。
将当前节点的Next节点的Next指针指向当前节点，然后将当前节点的Next指针置空。
返回新的头节点。
*/
func (node *Node) ReverseByRecursion() *Node {
	if node == nil || node.Next == nil {
		return node
	}
	newHead := node.Next.ReverseByRecursion()
	node.Next.Next = node
	node.Next = nil
	return newHead
}

func (node *Node) Print() {
	fmt.Printf("%d -> ", node.Val)

	if node.Next != nil {
		node = node.Next
		node.Print()
	}
}
