package main

// https://leetcode.cn/problems/fibonacci-number/
func fib(n int) int {
	mod := 1000000007
	f1 := 0
	f2 := 1
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	for n-1 > 0 {
		f1, f2 = f2%mod, (f1+f2)%mod
		n--
	}
	return f2 % mod
}
