package main

func groupThePeople(groupSizes []int) (ans [][]int) {
	// hashmap
	m := map[int]int{}
	for i, v := range groupSizes {
		if val, ok := m[v]; !ok || len(ans[val]) == v {
			// 开始新一波
			m[v] = len(ans)
			ans = append(ans, []int{i})
		} else {
			ans[val] = append(ans[val], i)
		}
	}
	return
}
