package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	// 创建一个新的栈
	stack := NewStack()

	// 压入元素
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	// 打印栈
	stack.Print() // 输出: [1 2 3]

	// 弹出元素
	fmt.Println(stack.Pop()) // 输出: 3

	// 查看栈顶元素
	fmt.Println(stack.Peek()) // 输出: 2

	// 打印栈的大小
	fmt.Println(stack.Size()) // 输出: 2

	// 判断栈是否为空
	fmt.Println(stack.IsEmpty()) // 输出: false
}
