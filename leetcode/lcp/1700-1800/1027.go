package main

// 1027 最长等差数列
// hash存储每个key
func longestArithSeqLength(nums []int) (ans int) {
	m := map[int]map[int]int{}
	for i := range nums {
		for j := 0; j < i; j++ {
			d := nums[i] - nums[j]
			if m[d] == nil {
				m[d] = map[int]int{}
			}
			if m[d][j] == 0 {
				m[d][i] = 2
				if ans < 2 {
					ans = 2
				}
			} else {
				m[d][i] = m[d][j] + 1
				if ans < m[d][i] {
					ans = m[d][i]
				}
			}
		}
	}
	return
}
