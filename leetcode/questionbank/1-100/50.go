package main

// 2 -2
func myPow(x float64, n int) (ans float64) {
	ans = 1
	if n < 0 {
		n = -n
		x = 1 / x
	}
	for n != 0 {
		if n&1 == 1 {
			ans *= x
		}
		x = x * x
		n >>= 1
	}
	return
}
