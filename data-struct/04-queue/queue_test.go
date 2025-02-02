package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	// 创建一个新的队列
	queue := NewQueue()

	// 入队元素
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// 打印队列
	queue.Print() // 输出: [1 2 3]

	// 出队元素
	fmt.Println(queue.Dequeue()) // 输出: 1

	// 查看队首元素
	fmt.Println(queue.Peek()) // 输出: 2

	// 打印队列的大小
	fmt.Println(queue.Size()) // 输出: 2

	// 判断队列是否为空
	fmt.Println(queue.IsEmpty()) // 输出: false
}
