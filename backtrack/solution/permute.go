package solution

// https://greyireland.gitbook.io/algorithm-pattern/suan-fa-si-wei/backtrack
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

func Permute(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	// 标记这个元素是否已经添加到结果集
	visited := make([]bool, len(nums))
	dfs3(nums, visited, list, &result)
	return result
}

func dfs3(nums []int, visited []bool, list []int, res *[][]int) {
	if len(list) == len(nums) {
		ans := make([]int, len(list))
		copy(ans, list)
		*res = append(*res, ans)
		return
	}
	for i := 0; i < len(nums); i++ {
		if visited[i] {
			continue
		}
		list = append(list, nums[i])
		visited[i] = true
		dfs3(nums, visited, list, res)
		visited[i] = false
		list = list[:len(list)-1]
	}
}
