package main

func numWays(n int) int {
	mod := 1000000007
	f1 := 1
	f2 := 1
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 1
	}
	if n == 2 {
		return 2
	}
	for n-1 > 0 {
		f1, f2 = f2%mod, (f1+f2)%mod
		n--
	}
	return f2 % mod
}
