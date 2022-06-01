package _022_05

// https://leetcode.cn/problems/Jf1JuT/

// 也可以使用dfs来找环
// visit[],0没有遍历，1遍历中，2完成遍历，如果dfs一个点的时候发现一个为1的点，则判定成环，return false
func AlienOrder(words []string) string {
	// 记录点
	node := make(map[byte]bool, 0)
	// 记录入度
	inL := make(map[byte]int)
	// 记录出度
	outL := make(map[byte][]byte)
	// 记下边
	edge := make(map[byte]byte)
	// 存node
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if !node[words[i][j]] {
				node[words[i][j]] = true
			}
		}
	}
	// 并查集存判断
	parent := make([]int, 26)
	for n := range node {
		parent[n-'a'] = int(n - 'a')
	}
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			// 判定 超出边界wer we
			if i+1 < len(words) && j == len(words[i+1]) {
				return ""
			}
			if i+1 < len(words) && words[i][j] != words[i+1][j] {
				// 判定 z，z
				if words[i+1][j] != words[i][j] && edge[words[i+1][j]] == words[i][j] {
					return ""
				}
				inL[words[i+1][j]] += 1
				outL[words[i][j]] = append(outL[words[i][j]], words[i+1][j])
				// 如果两个单词不一致
				if words[i][j] != words[i+1][j] {
					// 如果边里面顶点不匹配
					// 避免重复边 a->b, a->c, a->b
					if edge[words[i][j]] != words[i+1][j] {
						// z->a->b->z
						if ok, p := union(int(words[i][j]-'a'), int(words[i+1][j]-'a'), parent); !ok {
							if p == int(words[i][j]-'a') {
								return ""
							}
						}
					}
					// 会产生覆盖问题
					edge[words[i][j]] = words[i+1][j]
				}
				break
			}
		}
	}
	zeroIn := make([]byte, 0)
	for node := range node {
		if inL[node] == 0 {
			zeroIn = append(zeroIn, node)
		}
	}
	res := ""
	for len(zeroIn) != 0 {
		node := zeroIn[0]
		zeroIn = zeroIn[1:]
		res += string(node)
		for i := 0; i < len(outL[node]); i++ {
			inL[outL[node][i]] -= 1
			if inL[outL[node][i]] == 0 {
				zeroIn = append(zeroIn, outL[node][i])
			}
		}
	}
	return res
}

func find(t int, parent []int) int {
	for parent[t] != t {
		t = parent[t]
	}
	return t
}

func union(x, y int, parent []int) (bool, int) {
	xp := find(x, parent)
	yp := find(y, parent)
	if xp == yp {
		return false, yp
	}
	parent[xp] = yp //将x的根节点指向y的根节点
	return true, yp
}
