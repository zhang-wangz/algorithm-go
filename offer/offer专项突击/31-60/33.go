package _1_60

func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for i := 0; i < len(strs); i++ {
		cnt := [26]int{}
		for j := 0; j < len(strs[i]); j++ {
			cnt[strs[i][j]-'a']++
		}
		mp[cnt] = append(mp[cnt], strs[i])
	}
	ans := make([][]string, 0)
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
