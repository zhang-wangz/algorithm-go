package solution

import "sort"

// https://leetcode.cn/problems/subsets-ii/
// 给你一个整数数组nums，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
// 解集不能包含重复的子集。返回的解集中，子集可以按任意顺序排列。

func SubsetsWithDup(nums []int) [][]int {
	res := make([][]int, 0)
	lst := make([]int, 0)
	sort.Ints(nums)
	dfs2(lst, nums, 0, &res)
	return res
}

func dfs2(lst, nums []int, p int, res *[][]int) {
	ans := make([]int, len(lst))
	copy(ans, lst)
	*res = append(*res, ans)
	for i := p; i < len(nums); i++ {
		if i != p && nums[i] == nums[i-1] {
			continue
		}
		lst = append(lst, nums[i])
		dfs2(lst, nums, i+1, res)
		lst = lst[:len(lst)-1]
	}
}
