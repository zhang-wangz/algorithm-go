package main

// 1079
func numTilePossibilities(tiles string) (ans int) {
	nums := [26]int{}
	// 统计字母出现的次数
	for _, c := range tiles {
		nums[c-'A']++
	}
	var dfs func()
	dfs = func() {
		for i, v := range nums {
			if v > 0 {
				ans++
				nums[i]--
				dfs()
				nums[i]++
			}
		}
	}
	dfs()
	return
}
