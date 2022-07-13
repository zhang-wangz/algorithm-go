package unionFind

type UnionFind struct {
	parent []int
	cnt    []int
	groups int
	max    int
}

func initUF(n int) *UnionFind {
	// 初始时，总共有n个分组
	uf := &UnionFind{groups: n, max: 1}
	uf.parent = make([]int, n)
	uf.cnt = make([]int, n)
	for i := 0; i < n; i++ {
		// 初始时，每个元素单独一个分组
		uf.parent[i] = i
		// 每个分组里只有一个元素
		uf.cnt[i] = 1
	}
	return uf
}

// find 递归查找元素x的根结点，查找的同时将该组所有元素都直接指向根节点（路径压缩）
func (u *UnionFind) find(x int) int {
	if u.parent[x] != x {
		u.parent[x] = u.find(u.parent[x])
	}
	return u.parent[x]
}

// union 合并两个分组
func (u *UnionFind) union(x, y int) {
	xp := u.find(x)
	yp := u.find(y)
	if xp == yp {
		// 已经是同一个分组了，直接返回
		return
	}
	// 我们将小分组合并到大分组（这一步不是必须的）
	if u.cnt[yp] > u.cnt[xp] {
		xp, yp = yp, xp
	}
	// 大分组的元素数量增加
	u.cnt[xp] += u.cnt[yp]
	if u.cnt[xp] > u.max {
		u.max = u.cnt[xp]
	}
	// 小分组消失，让元素数量变0
	u.cnt[yp] = 0
	// 合并只需要小分组的根指向大分组任意一个元素即可
	// 这里我们让小分组的根指向大分组的根
	u.parent[yp] = xp
	// 总的分组数减少
	u.groups--
}
