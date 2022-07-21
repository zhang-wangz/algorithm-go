package main

func printNumbers(n int) (ans []int) {
	max := 0
	for n != 0 {
		n--
		max = max*10 + 9
	}
	for i := 1; i <= max; i++ {
		ans = append(ans, i)
	}
	return ans
}
