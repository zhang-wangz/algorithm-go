package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const MOD = int(1e9 + 7)
	var m [2005][2005]int
	m[1][0] = 1
	m[1][1] = 1
	for i := 2; i < 2003; i++ {
		m[i][0] = 1
		for j := 1; j <= i; j++ {
			m[i][j] = (m[i-1][j-1] + m[i-1][j]) % MOD
		}
	}
	var n, k int
	fmt.Fscan(in, &n, &k)
	hashMap := make(map[string]int)
	var s string
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &s)
		c := []byte(s)
		sort.Slice(c, func(i, j int) bool {
			return c[i] < c[j]
		})
		s = string(c)
		hashMap[s]++
	}
	mLen := len(hashMap)
	count := make([]int, mLen)
	var p int
	for _, v := range hashMap {
		count[p] = v
		p++
	}
	ans := make([][]int, mLen+1)
	for i := range ans {
		ans[i] = make([]int, k+1)
	}
	ans[0][0] = 1
	for i := 1; i <= mLen; i++ {
		for j := 0; j <= k; j++ {
			ans[i][j] = ans[i-1][j]
		}
		for b := 1; b <= count[i-1]; b++ {
			num := b * (b - 1) / 2
			if num > k {
				break
			}
			for j := num; j <= k; j++ {
				ans[i][j] = (ans[i][j] + ans[i-1][j-num]*m[count[i-1]][b]) % MOD
			}
		}

	}
	fmt.Fprintln(out, ans[mLen][k])
}
