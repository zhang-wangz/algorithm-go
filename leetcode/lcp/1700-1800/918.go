package main

// 918
func maxSubarraySumCircular(nums []int) int {
	total := 0
	curmax := 0
	curmin := 0
	maxsum := nums[0]
	minsum := nums[0]
	for _, a := range nums {
		curmax = max(curmax+a, a)
		maxsum = max(maxsum, curmax)
		curmin = min(curmin+a, a)
		minsum = min(minsum, curmin)
		total += a
	}
	if maxsum > 0 {
		return max(maxsum, total-minsum)
	} else {
		return maxsum
	}
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
//func min(a, b int) int {
//	if a < b {
//		return a
//	} else {
//		return b
//	}
//}
