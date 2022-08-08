package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m, q, k int
	fmt.Fscan(in, &n, &m, &q, &k)
	dst := make([]int, n)
	for i := range dst {
		dst[i] = n
	}
	var queue []int
	var u, v, cnt int
	for i := 0; i < q; i++ {
		fmt.Fscan(in, &u)
		u--
		dst[u] = 0
		queue = append(queue, u)
		cnt++
	}
	g := make([][]int, n)
	for i := 0; i < m; i++ {
		fmt.Fscan(in, &u, &v)
		u--
		v--
		g[u] = append(g[u], v)
		g[v] = append(g[v], u)
	}

	for day := 1; len(queue) > 0; day++ {
		for layer := 0; layer < k*day && len(queue) > 0; layer++ {
			floor := queue
			queue = []int{}
			for len(floor) > 0 {
				u = floor[0]
				floor = floor[1:]
				for _, cv := range g[u] {
					if dst[cv] == n {
						dst[cv] = day
						queue = append(queue, cv)
					}
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Fprintf(out, "%d ", dst[i])
	}

}
