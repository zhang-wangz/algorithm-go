package main

// https://leetcode.cn/problems/zhan-de-ya-ru-dan-chu-xu-lie-lcof/
func validateStackSequences(pushed []int, popped []int) bool {
	stack := make([]int, 0)
	cnt := 0
	if len(pushed) == 0 || len(popped) == 0 {
		return true
	}
	for i := 0; i < len(pushed); i++ {
		stack = append(stack, pushed[i])
		for len(stack) != 0 && stack[len(stack)-1] == popped[cnt] {
			stack = stack[:len(stack)-1]
			cnt++
			if cnt == len(popped) {
				return true
			}
		}
	}
	return false
}
