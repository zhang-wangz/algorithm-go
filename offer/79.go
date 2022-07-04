package main

// 选择元素，然后dfs，然后删除元素
//func subsets(nums []int) [][]int {
//
//}

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
