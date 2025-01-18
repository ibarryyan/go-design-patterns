package linked_list

import "testing"

func TestLinkedList(t *testing.T) {
	linkedList := NewNode(1, NewNode(2, nil))
	linkedList.Print()
}
