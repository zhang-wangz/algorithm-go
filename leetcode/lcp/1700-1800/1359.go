package main

// 1359 有效的快递序列数目
func countOrders(n int) int {
	const mod int = 1e9 + 7
	if n == 1 {
		return 1
	}
	ans := 1
	// 对i-1里面插入第i个p和d
	// P的位置有2i-1个，对应0，1，2，3，4...2n-2
	// 分别d有2n-1，2n-2，2n-3...1
	// 求和 (2n-1) * n
	for i := 2; i <= n; i++ {
		ans = ans * (i*2 - 1) % mod * i % mod
	}
	return ans
}
