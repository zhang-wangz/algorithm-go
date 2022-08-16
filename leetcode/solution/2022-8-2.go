package main

type MyCircularQueue struct {
	queue     []int
	top, tear int
	cap       int
}

func Constructor2(k int) MyCircularQueue {
	return MyCircularQueue{queue: make([]int, k+1), top: 0, tear: 0, cap: k}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	size := this.cap + 1
	this.queue[(this.tear+1)%size] = value
	this.tear += 1
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.top += 1
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	size := this.cap + 1
	return this.queue[(this.top+1)%size]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.tear]
}

func (this *MyCircularQueue) IsEmpty() bool {
	size := this.cap + 1
	this.tear = this.tear % size
	this.top = this.top % size
	return (this.tear)%size == this.top%size
}

func (this *MyCircularQueue) IsFull() bool {
	size := this.cap + 1
	this.tear = this.tear % size
	this.top = this.top % size
	return (this.tear+1)%size == this.top%size
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
