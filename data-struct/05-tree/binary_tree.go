package tree

import "fmt"

// TreeNode 定义二叉树节点结构体
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode 创建一个新的二叉树节点
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

// Insert 插入节点到二叉树中
func (root *TreeNode) Insert(val int) {
	if root == nil {
		return
	}
	if val < root.Val {
		if root.Left == nil {
			root.Left = NewTreeNode(val)
		} else {
			root.Left.Insert(val)
		}
	} else {
		if root.Right == nil {
			root.Right = NewTreeNode(val)
		} else {
			root.Right.Insert(val)
		}
	}
}

// Search 查找二叉树中是否存在某个值
func (root *TreeNode) Search(val int) bool {
	if root == nil {
		return false
	}
	if val == root.Val {
		return true
	} else if val < root.Val {
		return root.Left.Search(val)
	} else {
		return root.Right.Search(val)
	}
}

// Delete 删除二叉树中的某个节点
func (root *TreeNode) Delete(val int) *TreeNode {
	if root == nil {
		return nil
	}
	if val < root.Val {
		root.Left = root.Left.Delete(val)
	} else if val > root.Val {
		root.Right = root.Right.Delete(val)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		minNode := root.Right.FindMin()
		root.Val = minNode.Val
		root.Right = root.Right.Delete(minNode.Val)
	}
	return root
}

// FindMin 查找二叉树中的最小节点
func (root *TreeNode) FindMin() *TreeNode {
	if root.Left == nil {
		return root
	}
	return root.Left.FindMin()
}

// InorderTraversal 中序遍历二叉树
func (root *TreeNode) InorderTraversal() {
	if root == nil {
		return
	}
	root.Left.InorderTraversal()
	fmt.Printf("%d ", root.Val)
	root.Right.InorderTraversal()
}

// PreorderTraversal 前序遍历二叉树
func (root *TreeNode) PreorderTraversal() {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	root.Left.PreorderTraversal()
	root.Right.PreorderTraversal()
}

// PostorderTraversal 后序遍历二叉树
func (root *TreeNode) PostorderTraversal() {
	if root == nil {
		return
	}
	root.Left.PostorderTraversal()
	root.Right.PostorderTraversal()
	fmt.Printf("%d ", root.Val)
}
