package numbertheory

// 计算m的n次方对p取余的结果
func QuickPowandMod(m, n, p int) int {
	ans := 1
	for n >= 0 {
		if n&1 != 0 {
			ans = (ans * (m % p)) % p
		}
		n >>= 1
		m = ((m % p) * (m % p)) % p
	}
	return ans
}
