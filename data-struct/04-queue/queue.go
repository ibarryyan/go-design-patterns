package queue

import "fmt"

// Queue 定义队列结构体
type Queue struct {
	items []int
}

// NewQueue 创建一个新的队列
func NewQueue() *Queue {
	return &Queue{}
}

// Enqueue 将元素入队
func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

// Dequeue 出队并返回队首元素
func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		panic("Queue is empty")
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Peek 返回队首元素，但不出队
func (q *Queue) Peek() int {
	if q.IsEmpty() {
		panic("Queue is empty")
	}
	return q.items[0]
}

// IsEmpty 判断队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// Size 返回队列的大小
func (q *Queue) Size() int {
	return len(q.items)
}

// Print 打印队列中的所有元素
func (q *Queue) Print() {
	fmt.Println(q.items)
}
