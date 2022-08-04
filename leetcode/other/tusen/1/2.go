package main

import (
	"fmt"
)

const INF = 1<<63 - 1

type pair struct {
	v, w int
}

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	n2i := map[string]int{}
	id := 1
	w := make([][]pair, n+10)
	add := func(st string) int {
		if n2i[st] == 0 {
			n2i[st] = id
			id++
		}
		return n2i[st]
	}
	for i := 0; i < m; i++ {
		var n1, n2 string
		var we int
		fmt.Scanf("%s %s %d", &n1, &n2, &we)
		id1 := add(n1)
		id2 := add(n2)
		w[id1] = append(w[id1], pair{id2, we})
	}
	var k int
	fmt.Scanf("%d", &k)
	for i := 0; i < k; i++ {
		var n1, n2 string
		fmt.Scanf("%s %s", &n1, &n2)
		id1, id2 := add(n1), add(n2)
		dis := solve(id1, id2, w, n)
		if dis == INF {
			fmt.Println("INF")
		} else {
			fmt.Println(dis)
		}
	}
}

func solve(id1, id2 int, w [][]pair, n int) int {
	dis := make([]int, n+10)
	for i := range dis {
		dis[i] = INF
	}
	dis[id1] = 0
	q := []pair{{id1, 0}}
	for len(q) != 0 {
		x := q[0]
		q = q[1:]
		if x.w > dis[x.v] {
			continue
		}
		for _, v := range w[x.v] {
			node, we := v.v, v.w
			if dis[x.v]+we < dis[node] {
				dis[node] = dis[x.v] + we
				q = append(q, pair{node, dis[node]})
			}
		}
	}
	return dis[id2]
}
