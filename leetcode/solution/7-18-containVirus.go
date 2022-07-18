package solution

// https://leetcode.cn/problems/contain-virus/
// bfs搬砖题
type pair struct {
	i, j int
}

func containVirus(isInfected [][]int) int {
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	m, n := len(isInfected), len(isInfected[0])
	ans := 0
	for {
		neis := []map[pair]struct{}{}
		firewalls := make([]int, 0)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if isInfected[i][j] != 1 {
					continue
				}
				nei := map[pair]struct{}{}
				// idx从1开始，不然会和1，0搞混
				firewall, idx := 0, len(neis)+1
				q := []pair{{i, j}}
				isInfected[i][j] = -idx
				for len(q) != 0 {
					p := q[0]
					q = q[1:]
					for _, dir := range dirs {
						nx, ny := p.i+dir[0], p.j+dir[1]
						if nx >= 0 && nx < m && ny >= 0 && ny < n {
							if isInfected[nx][ny] == 1 {
								q = append(q, pair{nx, ny})
								isInfected[nx][ny] = -idx
								// 还存在负数的或者正数的不能被计算
							} else if isInfected[nx][ny] == 0 {
								firewall++
								// 必须是set，struct的默认比较规则是类型分别比较
								nei[pair{nx, ny}] = struct{}{}
							}
						}
					}
				}
				neis = append(neis, nei)
				firewalls = append(firewalls, firewall)
			}
		}
		// 如果没有被感染区域，就直接退出
		if len(neis) == 0 {
			break
		}
		// 确认最大firewalls数
		idx := 0
		for i := 1; i < len(neis); i++ {
			if len(neis[i]) > len(neis[idx]) {
				idx = i
			}
		}
		ans += firewalls[idx]
		// 恢复
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if isInfected[i][j] < 0 {
					if isInfected[i][j] != -idx-1 {
						// 其余恢复1，会被感染
						isInfected[i][j] = 1
					} else {
						// 该区域是有防火墙存在的，不被感染
						isInfected[i][j] = 2
					}
				}
			}
		}
		// 不是idx的都被感染
		for i, neighbor := range neis {
			if i != idx {
				for p := range neighbor {
					isInfected[p.i][p.j] = 1
				}
			}
		}
		if len(neis) == 1 {
			break
		}
	}
	return ans
}
