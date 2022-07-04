package _1_78

func mySqrt(x int) int {
	l, r := 0, x+1
	for l < r {
		mid := l + (r-l)>>1
		if mid*mid <= x {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l - 1
}
