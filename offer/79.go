package main

// 选择元素，然后dfs，然后删除元素
// 回溯主要是画决策树
func subsets(nums []int) (ans [][]int) {
	var dfs func(nums []int, startIdx int)
	set := []int{}
	dfs = func(nums []int, startIdx int) {
		ans = append(ans, append([]int(nil), set...))
		for i := startIdx; i < len(nums); i++ {
			set = append(set, nums[i])
			dfs(nums, i+1)
			set = set[:len(set)-1]
		}
	}
	dfs(nums, 0)
	return
}

// 状态压缩
func subsets2(nums []int) (ans [][]int) {
	n := len(nums)
	for mask := 0; mask < 1<<n; mask++ {
		set := []int{}
		for i, v := range nums {
			if mask>>i&1 == 1 {
				set = append(set, v)
			}
		}
		ans = append(ans, append([]int(nil), set...))
	}
	return
}
