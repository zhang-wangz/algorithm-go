package solution

import "sort"

// https://leetcode.cn/problems/permutations-ii/
// 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	list := make([]int, 0)
	visit := make([]bool, len(nums))
	sort.Ints(nums)
	dfs4(nums, list, visit, &res)
	return res
}

func dfs4(nums []int, list []int, visit []bool, res *[][]int) {
	if len(list) == len(nums) {
		ans := make([]int, len(list))
		copy(ans, list)
		*res = append(*res, ans)
	}
	for i := 0; i < len(nums); i++ {
		if i != 0 && nums[i] == nums[i-1] && !visit[i-1] {
			continue
		}
		list = append(list, nums[i])
		visit[i] = true
		dfs4(nums, list, visit, res)
		visit[i] = false
		list = list[:len(list)-1]
	}
}
