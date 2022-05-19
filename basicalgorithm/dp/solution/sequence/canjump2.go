package sequence

// https://leetcode.cn/problems/jump-game-ii/
// 给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
// 假设你总是可以到达数组的最后一个位置。

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	r := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		r[i] = i
		for j := 0; j < i; j++ {
			if nums[j]+j >= i {
				r[i] = min(r[j]+1, r[i])
			}
		}
	}
	return r[len(nums)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
