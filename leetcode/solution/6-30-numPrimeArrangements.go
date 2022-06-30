package solution

import "math"

// https://leetcode.cn/problems/prime-arrangements/

const N = 1000000007

func NumPrimeArrangements(n int) int {
	cnt := 0
	for i := 2; i <= n; i++ {
		if isNum(i) {
			cnt++
		}
	}
	return mul(cnt) % N * mul(n-cnt) % N
}
func mul(n int) int {
	num := 1
	for i := n; i >= 1; i-- {
		num = num * i % N
	}
	return num
}

func isNum(n int) bool {
	if n == 1 {
		return false
	}
	if n == 2 {
		return true
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
