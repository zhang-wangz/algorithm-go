package main

func validateStackSequences(pushed []int, popped []int) bool {
	st := make([]int, 0)
	i, j := 0, 0
	for j < len(popped) {
		if i == len(pushed) {
			return false
		}
		st = append(st, pushed[i])
		i++
		for len(st) > 0 && st[len(st)-1] == popped[j] {
			st = st[:len(st)-1]
			j++
		}
	}
	return true
}
