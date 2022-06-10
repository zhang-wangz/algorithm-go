package solution

// 给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 0; i <= num; i++ {
		res[0] = count(i)
	}
	return res
}

func count(num int) int {
	res := 0
	for num != 0 {
		num &= num - 1
		res++
	}
	return res
}
