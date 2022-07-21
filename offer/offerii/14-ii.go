package main

// 因为这时测试用例过大，有大数越界的问题，所以不能通过dp求解
// 需要通过数学方法求解
func cuttingRope2(n int) int {
	mod := 1000000007
	if n < 4 {
		return n - 1
	} else if n == 4 {
		return 4
	}
	a := n / 3
	b := n % 3
	if b == 0 {
		return pow(3, a) % mod
	} else if b == 1 {
		return pow(3, a-1) * 4 % mod
	} else if b == 2 {
		return pow(3, a) * 2 % mod
	}
	return 0
}

func pow(b, a int) int {
	mod := 1000000007
	res := 1
	for a != 0 {
		if a%2 == 1 {
			res *= b
			res %= mod
		}
		b = b * b % mod
		a >>= 1
	}
	return res
}
