package main

func missingNumber(nums []int) int {
	ans := 0
	for i := 0; i <= len(nums); i++ {
		ans ^= i
	}
	for _, num := range nums {
		ans ^= num
	}
	return ans
}
