package graph_105_119

import "strings"

// 需要考虑一下dfs
// 合理的typo
func alienOrder(words []string) string {
	graph := make(map[byte][]byte, 0)
	in := make(map[byte]int, 0)
	for _, v := range words[0] {
		in[byte(v)] = 0
	}
out:
	for i := 1; i < len(words); i++ {
		w1 := words[i-1]
		w2 := words[i]
		for _, v := range words[i] {
			// 防止把下面的关系修改了
			in[byte(v)] += 0
		}
		for j := 0; j < len(w1) && j < len(w2); j++ {
			if w1[j] != w2[j] {
				graph[w1[j]] = append(graph[w1[j]], w2[j])
				in[w2[j]]++
				continue out
			}
		}
		if len(w1) > len(w2) {
			return ""
		}
	}
	ans := strings.Builder{}
	zero := make([]byte, 0)
	for i, v := range in {
		if v == 0 {
			zero = append(zero, i)
		}
	}
	for len(zero) != 0 {
		node := zero[0]
		zero = zero[1:]
		ans.WriteByte(node)
		for _, v := range graph[node] {
			in[v]--
			if in[v] == 0 {
				zero = append(zero, v)
			}
		}
	}
	if ans.Len() == len(in) {
		return ans.String()
	}
	return ""
}
