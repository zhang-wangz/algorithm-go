package main

// 667 1-n的数字，成就k个不同的正数
// 思维题
// 前n-k个数据按照1，2,3,4...处理
// 后面的数据按照1，n，2，n-1。。。处理
// 相互结合即可
func constructArray(n int, k int) (ans []int) {
	for i := 1; i < n-k; i++ {
		ans = append(ans, i)
	}
	for i, j := n-k, n; i <= j; i++ {
		ans = append(ans, i)
		if i != j {
			ans = append(ans, j)
		}
		j--
	}
	return
}
