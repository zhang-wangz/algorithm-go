package solution

func rangeBitwiseAnd(left int, right int) int {
	var cnt int
	for left != right {
		left >>= 1
		right >>= 1
		cnt++
	}
	return left << cnt
}
