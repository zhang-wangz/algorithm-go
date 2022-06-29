package offer专项突击_1_30

func minSubArrayLen(target int, nums []int) int {
	left := 0
	right := 0
	sum := 0
	minV := 1<<31 - 1
	for right < len(nums) {
		t := nums[right]
		right++
		sum += t

		for sum >= target {
			minV = min(right-left, minV)
			t := nums[left]
			left++
			sum -= t
		}
	}
	if minV == 1<<31-1 {
		return 0
	}
	return minV
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
