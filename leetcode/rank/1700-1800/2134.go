package main

// 转化为统计0的个数，因为区间内为1，即区间内的0需要和外部进行交换，所以内部的0就是需要进行交换的次数
// 不断的进行滑动，最后所求的值就是我们所需要的
func minSwaps(nums []int) int {
	tot1 := 0
	for i := range nums {
		tot1 += nums[i]
	}
	cur1, max1 := 0, 0
	nums = append(nums, nums...)
	for i, num := range nums {
		cur1 += num
		if i >= tot1 {
			cur1 -= nums[i-tot1]
			if cur1 > max1 {
				max1 = cur1
			}
		}
	}
	return tot1 - max1
}
