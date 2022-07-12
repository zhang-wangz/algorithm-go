package main

// dfs
func findOrder2(numCourses int, prerequisites [][]int) []int {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses)
		result  []int
		valid   bool = true
		dfs     func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 {
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u)
	}

	for _, info := range prerequisites {
		// 记录入度边
		edges[info[1]] = append(edges[info[1]], info[0])
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return []int{}
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i]
	}
	return result
}

// typo // bfs在图里面
func findOrder(numCourses int, prerequisites [][]int) (ans []int) {
	zeroIn := make([]int, 0)
	in := make([]int, numCourses)
	graph := make([][]int, numCourses)
	for f := range graph {
		graph[f] = make([]int, 0)
	}
	for _, v := range prerequisites {
		in[v[0]]++
		graph[v[1]] = append(graph[v[1]], v[0])
	}
	for i, z := range in {
		if z == 0 {
			zeroIn = append(zeroIn, i)
		}
	}
	for len(zeroIn) != 0 {
		node := zeroIn[0]
		zeroIn = zeroIn[1:]
		ans = append(ans, node)
		// nei
		for _, v := range graph[node] {
			in[v]--
			if in[v] == 0 {
				zeroIn = append(zeroIn, v)
			}
		}
	}
	if len(ans) != numCourses {
		return []int{}
	}
	return
}
