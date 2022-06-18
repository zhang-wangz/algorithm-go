package swordReferoffer

// 状态压缩
func maxProduct(words []string) int {
	mask := make([]int, len(words))
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			mask[i] |= 1 << (words[i][j] - 'a')
		}
	}
	ans := 0
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words); j++ {
			if mask[i]&mask[j] == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
