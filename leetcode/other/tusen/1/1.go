package main

func removeElement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return j
}
