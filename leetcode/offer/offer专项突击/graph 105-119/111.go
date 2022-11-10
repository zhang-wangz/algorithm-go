package graph_105_119

// bfs
func calcEquation(equations [][]string, values []float64, queries [][]string) (ans []float64) {
	graph := [][]float64{}
	wordToId := map[string]int{}
	edge := [][]float64{}
	addWord := func(w string) int {
		id, has := wordToId[w]
		if !has {
			id = len(wordToId)
			wordToId[w] = id
			graph = append(graph, []float64{})
			edge = append(edge, []float64{})
		}
		return id
	}
	addEdge := func(w1 string, w2 string, v float64) {
		id1 := addWord(w1)
		id2 := addWord(w2)
		graph[id1] = append(graph[id1], float64(id2))
		graph[id2] = append(graph[id2], float64(id1))
		// 可以使用struct {to, weight}优化掉边的空间
		edge[id1] = append(edge[id1], v)
		edge[id2] = append(edge[id2], 1/v)
	}
	for i, v := range equations {
		addEdge(v[0], v[1], values[i])
	}

out:
	for _, quer := range queries {
		_, ok := wordToId[quer[0]]
		_, ok1 := wordToId[quer[1]]
		dist := make([]float64, len(wordToId))
		vis := make([]bool, len(wordToId))
		for f := range dist {
			dist[f] = 1.0
		}
		if !ok || !ok1 {
			ans = append(ans, -1)
			continue out
		}
		q := append([]int{}, wordToId[quer[0]])
		t := wordToId[quer[1]]
		for len(q) != 0 {
			node := q[0]
			q = q[1:]

			for i, v := range graph[node] {
				if !vis[int(v)] {
					dist[int(v)] = float64(dist[node]) * float64(edge[node][i])
					if int(v) == t {
						ans = append(ans, dist[int(v)])
						continue out
					}
					vis[int(v)] = true
					q = append(q, int(v))
				}
			}
		}
		ans = append(ans, -1)
	}
	return
}
