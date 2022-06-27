package offer

//给定一个非负整数 n ，请计算 0 到 n 之间的每个数字的二进制表示中 1 的个数，并输出一个数组。

func CountBits(n int) []int {
	ans := make([]int, 0)
	for i := 0; i <= n; i++ {
		cnt := 0
		num := i
		for num > 0 {
			num &= num - 1
			cnt++
		}
		ans = append(ans, cnt)
	}
	return ans
}
