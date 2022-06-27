package offer

func singleNumber(nums []int) int {
	ans := 0
	for i := 0; i < 64; i++ {
		sum := 0
		for j := 0; j < len(nums); j++ {
			sum += (nums[j] >> i) & 1
		}
		ans = ans ^ (sum%3)<<i
	}
	return ans
}
