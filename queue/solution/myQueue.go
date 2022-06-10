package solution

type MyQueue struct {
	stack []int
	back  []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		stack: []int{},
		back:  []int{},
	}
}
func (queue *MyQueue) push(x int) {
	if len(queue.back) != 0 {
		val := queue.back[len(queue.back)-1]
		queue.back = queue.back[:len(queue.back)-1]
		queue.stack = append(queue.stack, val)
	}
	queue.stack = append(queue.stack, x)
}

func (queue *MyQueue) pop() int {
	if len(queue.stack) != 0 {
		val := queue.stack[len(queue.stack)-1]
		queue.stack = queue.stack[:len(queue.stack)-1]
		queue.back = append(queue.back, val)
	}
	if len(queue.stack) == 0 {
		return 0
	}
	val := queue.back[len(queue.back)-1]
	queue.back = queue.back[:len(queue.back)-1]
	return val
}

func (queue *MyQueue) peek() int {
	if len(queue.stack) != 0 {
		val := queue.stack[len(queue.stack)-1]
		queue.stack = queue.stack[:len(queue.stack)-1]
		queue.back = append(queue.back, val)
	}
	if len(queue.stack) == 0 {
		return 0
	}
	val := queue.back[len(queue.back)-1]
	return val
}

func (queue *MyQueue) Empty() bool {
	return len(queue.back) == 0 && len(queue.stack) == 0
}
