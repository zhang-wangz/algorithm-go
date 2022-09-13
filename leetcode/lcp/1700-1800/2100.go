package main

// 2100
// 左边连续time天大于等于，右边小于等于
func goodDaysToRobBank(security []int, time int) (ans []int) {
	n := len(security)
	if time == 0 {
		for i := range security {
			ans = append(ans, i)
		}
		return
	}
	// 只记录了1到最后第2个
	left := make([]int, n)
	right := make([]int, n)
	sum := 0
	for i := 1; i < n; i++ {
		if security[i] <= security[i-1] {
			sum++
		} else {
			sum = 0
		}
		left[i] = sum
	}
	sum = 0
	for i := n - 2; i >= 0; i-- {
		if security[i] <= security[i+1] {
			sum++
		} else {
			sum = 0
		}
		right[i] = sum
	}
	for i := range left {
		if left[i] >= time && right[i] >= time {
			ans = append(ans, i)
		}
	}
	return
}
