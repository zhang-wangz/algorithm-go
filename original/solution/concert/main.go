package main

import (
	"fmt"
	"sort"
)

func main() {
	var m, n int
	fmt.Scanf("%d %d", &m, &n)
	data := make([][]byte, m)
	for i := 0; i < m; i++ {
		if data[i] == nil {
			data[i] = make([]byte, n)
		}
		fmt.Scanln(&data[i])
	}
	res := solve(data)
	fmt.Println(res)
}

func solve(memo [][]byte) int {
	m := len(memo)
	n := len(memo[0])
	data := make([][]int, m)
	treeNode := make([][]int, 0)
	idx := 0
	for i := 0; i < m; i++ {
		data[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if memo[i][j] == 'x' {
				treeNode = append(treeNode, []int{i, j})
				data[i][j] = idx
				idx++
			} else if memo[i][j] == '#' {
				data[i][j] = -2
			} else if memo[i][j] == '.' {
				data[i][j] = -1
			}
		}
	}
	if len(treeNode) == 1 {
		return 1
	}
	edge := make([][]int, 0)
	bfs := func(idx, fx, fy int) {
		vis := make([][]bool, m)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				vis[i] = make([]bool, n)
			}
		}
		q := make([][]int, 0)
		q = append(q, []int{fx, fy})
		vis[fx][fy] = true
		step := 0
		for len(q) != 0 {
			size := len(q)
			for i := 0; i < size; i++ {
				node := q[0]
				q = q[1:]
				if data[node[0]][node[1]] != -2 && data[node[0]][node[1]] != -1 && data[node[0]][node[1]] != idx {
					tidx := data[node[0]][node[1]]
					edge = append(edge, []int{idx, tidx, step})
				}
				dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
				for j := 0; j < len(dir); j++ {
					nx := node[0] + dir[j][0]
					ny := node[1] + dir[j][1]
					if nx >= 0 && nx < m && ny >= 0 && ny < n {
						if !vis[nx][ny] && data[nx][ny] != -2 {
							vis[nx][ny] = true
							q = append(q, []int{nx, ny})
						}
					}
				}
			}
			step++
		}
	}

	for i := 0; i < len(treeNode)-1; i++ {
		bfs(i, treeNode[i][0], treeNode[i][1])
	}
	sort.Slice(edge, func(i, j int) bool {
		return edge[i][2] < edge[j][2]
	})
	sum := 0
	parent := make([]int, len(treeNode))
	for i := 0; i < len(treeNode); i++ {
		parent[i] = -1
	}
	vis := map[int]bool{}
	for i := 0; i < len(edge); i++ {
		if union(edge[i][0], edge[i][1], parent) {
			sum += edge[i][2]
			vis[edge[i][0]] = true
			vis[edge[i][1]] = true
		}
		if len(vis) == len(edge)+1 {
			break
		}
	}
	for i := 0; i < len(treeNode); i++ {
		if findRoot(0, parent) != findRoot(i, parent) {
			return -1
		}
	}
	return sum + 1
}

// 查找节点所在的集合（查找根节点）
func findRoot(t int, parent []int) int {
	tmp := t
	for parent[tmp] != -1 {
		tmp = parent[tmp]
	}
	return tmp
}

// 将两个节点所在的树（集合）进行合并; 合并失败返回false，成功返回true
// 同时也可以检查两个节点是否在同一个集合
func union(x, y int, parent []int) bool {
	xroot := findRoot(x, parent)
	yroot := findRoot(y, parent)
	if xroot == yroot {
		return false
	}
	parent[xroot] = yroot //将x的根节点指向y的根节点
	return true
}
