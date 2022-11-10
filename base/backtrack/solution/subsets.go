package solution

// https://leetcode.cn/problems/subsets/

func Subsets(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) == 0 {
		return res
	}
	lst := make([]int, 0)
	dfs(lst, nums, 0, &res)
	return res
}

func dfs(lst, nums []int, p int, res *[][]int) {
	*res = append(*res, lst)
	for i := p; i < len(nums); i++ {
		lst = append(lst, nums[i])
		dfs(lst, nums, i+1, res)
		lst = lst[:len(lst)-1]
	}
}
