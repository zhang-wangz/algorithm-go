package offer专项突击_1_30

import "math"

// 负数的范围比正数大，所以采用都转换成负数进行求解

func Divide(a int, b int) int {
	if a == 0 || b == 1 {
		return a
	}
	// -1<<31
	minv := math.MinInt32
	// 1<<31-1
	maxv := math.MaxInt32
	minLimit := minv >> 1
	if a == minv && b == -1 {
		return maxv
	}
	f := false
	if a > 0 && b < 0 || a < 0 && b > 0 {
		f = true
	}
	ans := 0

	for a <= b {
		// 除是b，倍商是1
		d, c := b, 1
		for d <= minLimit && d+d >= a {
			d <<= 1
			c <<= 1
		}
		a -= d
		ans += c
	}
	if f {
		return -ans
	} else {
		return ans
	}
}

func quickadd(m, n int) int {
	ans := 0
	for m > 0 {
		if m&1 == 1 {
			ans = ans + n
		}
		m >>= 1
		n = n + n
	}
	return ans
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
