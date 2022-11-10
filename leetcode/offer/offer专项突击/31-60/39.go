package _1_60

import "math"

func largestRectangleArea(heights []int) int {
	left := make([]int, len(heights))
	right := make([]int, len(heights))
	stack := make([]int, 0)
	for i := 0; i < len(heights); i++ {
		right[i] = len(heights)
	}
	for i := 0; i < len(heights); i++ {
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			if heights[top] > heights[i] {
				stack = stack[:len(stack)-1]
				right[top] = i
			} else {
				break
			}
		}
		if len(stack) != 0 {
			left[i] = stack[len(stack)-1]
		} else {
			left[i] = -1
		}
		stack = append(stack, i)
	}
	ans := 0
	for i := 0; i < len(heights); i++ {
		ans = int(math.Max(float64((right[i]-left[i]-1)*heights[i]), float64(ans)))
	}
	return ans
}
