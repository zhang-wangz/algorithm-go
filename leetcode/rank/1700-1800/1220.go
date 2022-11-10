package main

// 1220  统计元音字母序列的数目
func countVowelPermutation(n int) (ans int) {
	const mod = 1e9 + 7
	f := make([][]int, n)
	f[0] = []int{1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		f[i] = make([]int, 5)
		f[i][0] = f[i-1][1]
		f[i][1] = f[i-1][0] + f[i-1][2]
		f[i][2] = (f[i-1][0] + f[i-1][1]) + f[i-1][3] + f[i-1][4]
		f[i][3] = f[i-1][2] + f[i-1][4]
		f[i][4] = f[i-1][0]
		for j := range f[i] {
			f[i][j] %= mod
		}
	}
	for i := 0; i < 5; i++ {
		ans += f[n-1][i] % mod
	}
	return ans % mod

}
