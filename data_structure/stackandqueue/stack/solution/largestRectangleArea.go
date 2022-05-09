package solution

// 给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
// 求在该柱状图中，能够勾勒出来的矩形的最大面积。

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	stack := make([]int, 0)
	left, right := make([]int, len(heights)), make([]int, len(heights))
	res := 0
	for i := 0; i < len(heights); i++ {
		right[i] = len(heights)
	}
	for i := 0; i < len(heights); i++ {
		for len(stack) != 0 && heights[i] <= heights[stack[len(stack)-1]] {
			right[stack[len(stack)-1]] = i
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			left[i] = stack[len(stack)-1]
		} else {
			left[i] = -1
		}
		stack = append(stack, i)
	}

	for j := 0; j < len(heights); j++ {
		res = max(res, (right[j]-left[j]-1)*heights[j])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
