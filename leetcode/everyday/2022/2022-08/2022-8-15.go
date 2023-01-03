package _022_08

// https://leetcode.cn/problems/design-circular-deque/
// 循环双端队列
type MyCircularDeque struct {
	front, rear int
	elements    []int
}

func Constructor3(k int) MyCircularDeque {
	return MyCircularDeque{elements: make([]int, k+1)}
}

func (q *MyCircularDeque) InsertFront(value int) bool {
	if q.IsFull() {
		return false
	}
	q.front = (q.front + 1) % len(q.elements)
	q.elements[q.front] = value
	return true
}

func (q *MyCircularDeque) InsertLast(value int) bool {
	if q.IsFull() {
		return false
	}
	// 尾部是直接插入
	// 所以获取的时候需要+1
	q.elements[q.rear] = value
	q.rear = (q.rear - 1 + len(q.elements)) % len(q.elements)
	return true
}

func (q *MyCircularDeque) DeleteFront() bool {
	if q.IsEmpty() {
		return false
	}
	q.front = (q.front - 1 + len(q.elements)) % len(q.elements)
	return true
}

func (q *MyCircularDeque) DeleteLast() bool {
	if q.IsEmpty() {
		return false
	}
	q.rear = (q.rear + 1) % len(q.elements)
	return true
}

func (q MyCircularDeque) GetFront() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elements[q.front]
}

func (q MyCircularDeque) GetRear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elements[(q.rear+1)%len(q.elements)]
}

func (q MyCircularDeque) IsEmpty() bool {
	return q.rear == q.front
}

func (q MyCircularDeque) IsFull() bool {
	return (q.rear-1+len(q.elements))%len(q.elements) == q.front
}
