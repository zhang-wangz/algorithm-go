package main

// 1954
// 求和计算公式 apple = n*(n+1)*2*(2n+1)
func minimumPerimeter(neededApples int64) int64 {
	d := 1
	var sum int64 = 12
	var cnt0 int64 = 12
	for sum < neededApples {
		// 当前层等于上一层+这一层附加的值
		cnt1 := cnt0 + int64(d*2+1)*4 + 4*(int64(d)*2) + int64(d+1)*2*4
		sum += cnt1
		cnt0 = cnt1
		d++
	}
	return int64(d) * 2 * 4
}
