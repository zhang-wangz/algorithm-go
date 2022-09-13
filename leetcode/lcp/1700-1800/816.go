package main

// 816. 模糊坐标
// 暴力遍历所有的可能性
// bf
func ambiguousCoordinates(S string) (ans []string) {
	s := S[1 : len(S)-1]
	n := len(s)
	for i := 0; i < n-1; i++ {
		for _, l := range build(s[:i+1]) {
			for _, r := range build(s[i+1:]) {
				ans = append(ans, "("+l+","+r+")")
			}
		}
	}
	return
}

func build(s string) (ans []string) {
	n := len(s)
	for i := 0; i < n; i++ {
		conn := "."
		if i == n-1 {
			conn = ""
		}
		s1, s2 := s[:i+1], s[i+1:]
		if !(len(s1) >= 2 && s1[0] == '0' || len(s2) > 0 && s2[len(s2)-1] == '0') {
			ans = append(ans, s1+conn+s2)
		}
	}
	return
}
