package main

func countEven(num int) (ans int) {
	y, x := num/10, num%10
	ans += y * 5
	ysum := 0
	for ; y > 0; y /= 10 {
		ysum += y % 10
	}
	// 需要减去 0，因为题目只需要正整数
	if ysum%2 == 1 {
		ans += (x+1)/2 - 1
	} else {
		ans += x/2 + 1 - 1
	}
	return
}
