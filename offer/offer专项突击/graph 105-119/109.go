package graph_105_119

func openLock(deadends []string, target string) int {
	const start = "0000"
	if target == start {
		return 0
	}
	dead := map[string]bool{}
	for _, w := range deadends {
		dead[w] = true
	}
	if dead[target] || dead[start] {
		return -1
	}
	get := func(word string) (ans []string) {
		s := []byte(word)
		for i, b := range s {
			s[i] = b - 1
			if s[i] < '0' {
				s[i] = '9'
			}
			ans = append(ans, string(s))
			s[i] = b + 1
			if s[i] > '9' {
				s[i] = '0'
			}
			ans = append(ans, string(s))
			s[i] = b
		}
		return
	}
	q := append(make([]string, 0), start)
	vis := map[string]bool{start: true}
	level := 0
	for len(q) != 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			node := q[0]
			q = q[1:]
			next := get(node)
			for _, nex := range next {
				if !vis[nex] && !dead[nex] {
					if nex == target {
						return level + 1
					}
					vis[nex] = true
					q = append(q, nex)
				}
			}
		}
		level++
	}
	return -1
}
