package numbertheory

// 计算m的n次方对p取余的结果
// 乘方快速幂
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

// 如果有多组数据加入时就可以使用快速幂进行求解

// 方阵快速幂
// 两个矩阵相乘，n*m,m*p
func matricmul(nums1, nums2 [][]int) [][]int {
	n := len(nums1)
	m := len(nums2)
	p := len(nums2[0])
	nums3 := make([][]int, n)
	for i := 0; i < len(nums3); i++ {
		nums3[i] = make([]int, p)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < p; j++ {
			sum := 0
			for k := 0; k < m; k++ {
				sum += nums1[i][k] * nums2[k][j]
			}
			nums3[i][j] = sum
		}
	}
	return nums3
}

func matrixQucikMod() {

}
