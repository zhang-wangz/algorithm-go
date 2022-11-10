package main

type pair struct {
	i, j int
}

func movingCount(m int, n int, k int) int {
	count := func(i, j int) bool {
		sum := 0
		for i != 0 {
			sum += i % 10
			i /= 10
		}
		for j != 0 {
			sum += j % 10
			j /= 10
		}
		return sum <= k
	}
	size := 0
	q := []*pair{}
	if count(0, 0) {
		q = append(q, &pair{0, 0})
		size++
	}
	vis := map[pair]bool{}
	vis[pair{0, 0}] = true
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		dirs := [][]int{{0, 1}, {1, 0}}
		for _, dir := range dirs {
			nx, ny := node.i+dir[0], node.j+dir[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n && !vis[pair{nx, ny}] && count(nx, ny) {
				q = append(q, &pair{nx, ny})
				vis[pair{nx, ny}] = true
				size++
			}
		}
	}
	return size
}
