package tusen2

// 把每一个字符转换成 *
const INF = 0x3f3f3f3f

func LadderLength(beginWord string, endWord string, wordList []string) int {
	g := make(map[string][]string, len(wordList)*len(endWord)+10)
	dis := make(map[string]int, len(wordList)*len(endWord)+10)
	for i := range wordList {
		w := wordList[i]
		b := []byte(w)
		for j, c := range b {
			b[j] = '*'
			g[w] = append(g[w], string(b))
			g[string(b)] = append(g[string(b)], w)
			dis[w] = INF
			dis[string(b)] = INF
			b[j] = c
		}
	}
	// 初始化beginWord
	if _, ok := g[beginWord]; !ok {
		w := beginWord
		b := []byte(w)
		for j, c := range beginWord {
			b[j] = '*'
			g[w] = append(g[w], string(b))
			g[string(b)] = append(g[string(b)], w)
			b[j] = byte(c)
		}
	}
	if _, ok := g[endWord]; !ok {
		return 0
	}
	dis[beginWord] = 0
	q := []string{beginWord}
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		if node == endWord {
			return dis[node]/2 + 1
		}
		for _, v := range g[node] {
			if dis[v] == INF {
				dis[v] = dis[node] + 1
				q = append(q, v)
			}
		}
	}
	return 0
}
