package solution

// https://leetcode.cn/problems/fibonacci-number/
func fib(N int) int {
	return dfs(N)
}

var m map[int]int = make(map[int]int)

func dfs(n int) int {
	if n < 2 {
		return n
	}
	// 读取缓存
	if m[n] != 0 {
		return m[n]
	}
	ans := dfs(n-2) + dfs(n-1)
	// 缓存已经计算过的值
	m[n] = ans
	return ans
}
