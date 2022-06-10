package solution

type MinStack struct {
	Min   []int
	Stack []int
}

func (minStack *MinStack) Constructor() *MinStack {
	return &MinStack{
		Min:   make([]int, 0),
		Stack: make([]int, 0),
	}
}

func (minStack *MinStack) Top() int {
	if len(minStack.Stack) == 0 {
		return 0
	}
	return minStack.Stack[len(minStack.Stack)-1]
}

func (minStack *MinStack) Pop() {
	if len(minStack.Stack) == 0 {
		return
	}
	minStack.Stack = minStack.Stack[:len(minStack.Stack)-1]
	minStack.Min = minStack.Min[:len(minStack.Min)-1]
}

func (minStack *MinStack) GetMin() int {
	if len(minStack.Min) == 0 {
		return 1 << 31
	} else {
		return minStack.Min[len(minStack.Min)-1]
	}
}

func (minStack *MinStack) Push(x int) {
	min := minStack.GetMin()
	if x < min {
		minStack.Min = append(minStack.Min, x)
	} else {
		minStack.Min = append(minStack.Min, min)
	}
	minStack.Stack = append(minStack.Stack, x)
}
