package main

// 分成两半，简单来说就是奇偶分
func maxDepthAfterSplit(seq string) (ans []int) {
	d := 0
	for _, c := range seq {
		if c == '(' {
			ans = append(ans, d%2)
			d++
		} else {
			d--
			ans = append(ans, d%2)
		}
	}
	return ans
}
