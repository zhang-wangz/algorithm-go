package _1_60

import "math"

// 给定一个由 0 和 1 组成的矩阵 matrix ，找出只包含 1 的最大矩形，并返回其面积。
//
//注意：此题 matrix 输入格式为一维 01 字符串数组。

func MaximalRectangle(matrix []string) int {
	if len(matrix) == 0 {
		return 0
	}
	ans := 0
	height := make([]int, len(matrix[0]))
	for _, v := range matrix {
		for i := 0; i < len(v); i++ {
			if v[i] == '1' {
				height[i] = height[i] + 1
			} else {
				height[i] = 0
			}
		}
		res := maxArea(height)
		ans = int(math.Max(float64(ans), float64(res)))
	}
	return ans
}

func maxArea(heights []int) int {
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
