package solution

import "sort"

// https://leetcode.cn/problems/cut-off-trees-for-golf-event/
// 你被请来给一个要举办高尔夫比赛的树林砍树。树林由一个 m x n 的矩阵表示， 在这个矩阵中：.
// 0 表示障碍，无法触碰
// 1 表示地面，可以行走
// 比 1 大的数 表示有树的单元格，可以行走，数值表示树的高度
// 每一步，你都可以向上、下、左、右四个方向之一移动一个单位，如果你站的地方有一棵树，那么你可以决定是否要砍倒它。
// 你需要按照树的高度从低向高砍掉所有的树，每砍过一颗树，该单元格的值变为 1（即变为地面）。
// 你将从 (0, 0) 点开始工作，返回你砍完所有树需要走的最小步数。 如果你无法砍完所有的树，返回 -1 。
// 可以保证的是，没有两棵树的高度是相同的，并且你至少需要砍倒一棵树。
func CutOffTree(forest [][]int) int {
	m := len(forest)
	n := len(forest[0])
	tree := make([][]int, 0)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			//  是树的高度进行统计
			if forest[i][j] > 1 {
				tree = append(tree, []int{i, j})
			}
		}
	}
	// 按照高度排列
	sort.Slice(tree, func(i, j int) bool {
		return forest[tree[i][0]][tree[i][1]] < forest[tree[j][0]][tree[j][1]]
	})
	bfs := func(fx, fy, dx, dy int) int {
		if fx == dx && fy == dy {
			return 0
		}
		vis := make([][]bool, m)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				vis[i] = make([]bool, n)
			}
		}
		q := make([][]int, 0)
		q = append(q, []int{fx, fy})
		vis[fx][fy] = true
		step := 1
		for len(q) != 0 {
			size := len(q)
			for i := 0; i < size; i++ {
				node := q[0]
				q = q[1:]
				dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
				for j := 0; j < len(dir); j++ {
					nx := node[0] + dir[j][0]
					ny := node[1] + dir[j][1]
					if nx >= 0 && nx < m && ny >= 0 && ny < n {
						if !vis[nx][ny] && forest[nx][ny] > 0 {
							if nx == dx && ny == dy {
								return step
							}
							vis[nx][ny] = true
							q = append(q, []int{nx, ny})
						}
					}
				}
			}
			step++
		}
		return -1
	}
	fx, fy := 0, 0
	sum := 0
	for i := 0; i < len(tree); i++ {
		step := bfs(fx, fy, tree[i][0], tree[i][1])
		if step == -1 {
			return -1
		}
		sum += step
		fx = tree[i][0]
		fy = tree[i][1]
	}
	return sum
}
