package tree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	// 创建一个新的二叉树
	root := NewTreeNode(5)
	root.Insert(3)
	root.Insert(7)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	// 中序遍历二叉树
	fmt.Println("Inorder Traversal:")
	root.InorderTraversal() // 输出: 2 3 4 5 6 7 8

	// 前序遍历二叉树
	fmt.Println("\nPreorder Traversal:")
	root.PreorderTraversal() // 输出: 5 3 2 4 7 6 8

	// 后序遍历二叉树
	fmt.Println("\nPostorder Traversal:")
	root.PostorderTraversal() // 输出: 2 4 3 6 8 7 5

	// 查找节点
	fmt.Println("\nSearch 4:", root.Search(4)) // 输出: true

	// 删除节点
	root.Delete(3)
	fmt.Println("\nAfter Deleting 3:")
	root.InorderTraversal() // 输出: 2 4 5 6 7 8
}
