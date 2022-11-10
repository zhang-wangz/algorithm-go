package challenge

import "sort"

// https://leetcode.cn/problems/combination-sum/

func CombinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	list := make([]int, 0)
	sum := 0
	sort.Ints(candidates)
	dfs(&res, list, sum, target, candidates, 0)
	return res
}

func dfs(res *[][]int, list []int, sum, target int, nums []int, p int) {
	if sum == target {
		ans := make([]int, len(list))
		copy(ans, list)
		*res = append(*res, ans)
	}
	if sum > target {
		return
	}
	for i := p; i < len(nums); i++ {
		sum += nums[i]
		if sum > target {
			break
		}
		list = append(list, nums[i])
		dfs(res, list, sum, target, nums, i)
		list = list[:len(list)-1]
		sum -= nums[i]
	}
}
