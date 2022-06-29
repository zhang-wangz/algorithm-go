package offer

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0)
	ans := make([]int, len(temperatures))
	right := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for len(stack) != 0 {
			top := temperatures[stack[len(stack)-1]]
			if top < temperatures[i] {
				stack = stack[:len(stack)-1]
				right[top] = i
			} else {
				break
			}
		}
		stack = append(stack, i)
	}
	for i := 0; i < len(temperatures); i++ {
		if right[i] != 0 {
			ans[i] = right[i] - i
		} else {
			ans[i] = 0
		}
	}
	return ans
}
