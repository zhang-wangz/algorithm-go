package _022_07

func largestComponentSize(nums []int) int {
	m := 0
	for _, num := range nums {
		m = max(m, num)
	}
	parent := make([]int, m+1)
	rank := make([]int, m+1)
	for i := 0; i < len(parent); i++ {
		parent[i] = i
	}
	var find func(x int) int
	find = func(x int) int {
		if parent[x] != x {
			x = find(parent[x])
		}
		return x
	}
	union := func(x, y int) {
		x, y = find(x), find(y)
		if x == y {
			return
		}
		if rank[x] > rank[y] {
			parent[y] = x
		} else if rank[x] < rank[y] {
			parent[x] = y
		} else {
			parent[y] = x
			rank[x]++
		}
	}

	for _, num := range nums {
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				union(num, i)
				union(num, num/i)
			}
		}
	}
	ct := map[int]int{}
	ans := 0
	for i := 0; i < len(nums); i++ {
		x := find(nums[i])
		ct[x]++
		ans = max(ans, ct[x])
	}
	return ans
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
