package stack

import "fmt"

// Stack 定义栈结构体
type Stack struct {
	items []int
}

// NewStack 创建一个新的栈
func NewStack() *Stack {
	return &Stack{}
}

// Push 将元素压入栈顶
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop 弹出栈顶元素
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item
}

// Peek 返回栈顶元素，但不弹出
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		panic("Stack is empty")
	}
	return s.items[len(s.items)-1]
}

// IsEmpty 判断栈是否为空
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

// Size 返回栈的大小
func (s *Stack) Size() int {
	return len(s.items)
}

// Print 打印栈中的所有元素
func (s *Stack) Print() {
	fmt.Println(s.items)
}
