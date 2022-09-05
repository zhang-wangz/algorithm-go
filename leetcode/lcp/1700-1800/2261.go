package main

// 2261 含最多 K 个可整除元素的子数组
func countDistinct(nums []int, k int, p int) (ans int) {
	set := map[[200]int]struct{}{}
	for i := range nums {
		arr, cnt, idx := [200]int{}, 0, 0
		for _, v := range nums[i:] {
			if v%p == 0 {
				if cnt++; cnt > k {
					break
				}
			}
			arr[idx] = v
			idx++
			set[arr] = struct{}{}
		}
	}
	return len(set)
}
