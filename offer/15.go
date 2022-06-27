package offer

// https://leetcode.cn/problems/VabMRr/
func findAnagrams(s string, p string) []int {
	need := [26]int{}
	win := [26]int{}
	ans := make([]int, 0)
	if len(p) > len(s) {
		return ans
	}
	for i := 0; i < len(p); i++ {
		need[p[i]-'a']++
		win[s[i]-'a']++
	}
	if need == win {
		ans = append(ans, 0)
	}
	// 维护固定长度len(pattern)
	for i := len(p); i < len(s); i++ {
		win[s[i]-'a']++
		win[s[i-len(p)]-'a']--
		if need == win {
			ans = append(ans, i-len(p)+1)
		}
	}
	return ans
}
