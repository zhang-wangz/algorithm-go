package sequence

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶
// https://leetcode.cn/problems/climbing-stairs/
func climbStairs(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	f := make([]int, n+1)
	f[1] = 1
	f[2] = 2
	for i := 3; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
