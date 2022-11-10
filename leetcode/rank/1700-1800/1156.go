package main

// 贪心
func maxRepOpt1(text string) (ans int) {
	n := len(text)
	cnt := [26]int{}
	for _, c := range text {
		cnt[c-'a']++
	}
	for i := 0; i < n; {
		le := 0
		for i+le < n && text[i+le] == text[i] {
			le++
		}
		j := i + le + 1
		w := 0
		for j+w < n && text[j+w] == text[i] {
			w++
		}
		ans = max(ans, min(le+w+1, cnt[text[i]-'a']))
	}
	return
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	} else {
//		return b
//	}
//}
//func min(a, b int)int { if a < b{return a}else {return b}}
