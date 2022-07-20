package main

//https://leetcode.cn/problems/coin-bonus/
// 发leetcoin
func bonus(n int, leadership [][]int, operations [][]int) []int {
	begin, end := make([]int, n+1), make([]int, n+1)
	MOD := 1000000007
	id := 1
	g := make([][]int, n+1)
	for i := 0; i < len(g); i++ {
		g[i] = make([]int, 0)
	}
	ans := make([]int, 0)
	for _, l := range leadership {
		g[l[0]] = append(g[l[0]], l[1])
	}
	var dfs func(idx int)
	dfs = func(idx int) {
		begin[idx] = id
		for _, v := range g[idx] {
			dfs(v)
		}
		end[idx] = id
		id++
	}
	dfs(1)
	// 区间查询 + 区间更新 ： 前置需要了解差分+树状数组
	// 一个存储差分数组的前缀和，一个存储i*b[i]的变化
	// 因为需要求a[i]前缀和的增量变化，同时原来arr为空，所以sum可以默认全为0不需要考虑
	// sum[r] + (r+1)* ask1[r] - ask2[r]
	// sum[l-1] + (l-1+1)*ask1[l-1] - ask[l-1]
	trb := make([]int, n+1)
	trib := make([]int, n+1)
	add := func(tr []int, idx, k int) {
		for i := idx; i <= n; i += i & -i {
			tr[i] += k
		}
	}
	ask := func(tr []int, idx int) int {
		ans := 0
		for i := idx; i >= 1; i -= i & -i {
			ans += tr[i]
		}
		return ans
	}
	update := func(l, r int, k int) {
		add(trb, l, k)
		add(trb, r+1, -k)
		add(trib, l, l*k)
		add(trib, r+1, -(r+1)*k)
	}
	asklr := func(l, r int) int {
		return (r+1)*ask(trb, r) - ask(trib, r) - (l*ask(trb, l-1) - ask(trib, l-1))
	}
	for _, opr := range operations {
		if opr[0] == 1 {
			update(end[opr[1]], end[opr[1]], opr[2])
		} else if opr[0] == 2 {
			update(begin[opr[1]], end[opr[1]], opr[2])
		} else {
			ak := asklr(begin[opr[1]], end[opr[1]]) % MOD
			ans = append(ans, ak)
		}
	}
	return ans
}
