package main

func grayCode(n int) []int {
	ans := make([]int, 1<<n)
	for i := range ans {
		ans[i] = i/2 ^ i
	}
	return ans
}
