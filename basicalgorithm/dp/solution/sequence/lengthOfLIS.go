package sequence

// 给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。

// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
// 例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	r := make([]int, len(nums)+1)
	maxN := 1
	for i := 0; i < len(nums); i++ {
		r[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				r[i] = max(r[i], r[j]+1)
				if r[i] > maxN {
					maxN = r[i]
				}
			}
		}
	}
	return maxN
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
