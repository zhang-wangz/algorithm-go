package main
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i, num := range nums {
		if num <= 0 {
			nums[i] = n + 1
		}
	}
	for _, num := range nums {
		num = abs(num)
		if num <= n {
			nums[num - 1] = -abs(nums[num-1])
		}
	}
	for i, num := range nums {
		if num > 0 {
			return  i + 1
		}
	}
	return n + 1
}

func abs(x int)int {
	if x < 0 {
		return -x
	}else {
		return x
	}
}