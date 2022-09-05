package main

// 2202 K 次操作后最大化顶端元素
func maximumTop(nums []int, k int) (ans int) {
	n := len(nums)
	if n == 1 || k == 0 {
		if k%2 == 1 {
			return -1
		}
		return nums[0]
	}
	if k < n {
		nums = append(nums[:k-1], nums[k])
	} else if k == n {
		nums = nums[:n-1]
	}
	for _, v := range nums {
		if v > ans {
			ans = v
		}
	}
	return
}
