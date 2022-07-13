package main

// 并查
// 最长连续序列且未排序
// 是 连续 不是递增
func longestConsecutive(nums []int) int {
	parent := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}
	union := func(x1, x2 int) {
		x, y := find(x1), find(x2)
		parent[x] = y
	}
	mp := map[int]int{}
	for i, num := range nums {
		if _, ok := mp[num]; !ok {
			mp[num] = i
			if id1, has := mp[num-1]; has {
				union(i, id1)
			}
			if id2, has := mp[num+1]; has {
				union(i, id2)
			}
		}
	}
	ans := -1
	cnt := map[int]int{}
	for _, v := range parent {
		v = find(v)
		cnt[v]++
		if cnt[v] > ans {
			ans = cnt[v]
		}
	}
	return ans
}
