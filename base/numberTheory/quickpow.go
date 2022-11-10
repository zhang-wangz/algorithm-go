package numbertheory

// 计算m的n次方对p取余的结果
// 乘方快速幂
func QuickPowandMod(m, n, p int) int {
	ans := 1
	for n > 0 {
		if n&1 != 0 {
			ans = (ans * (m % p)) % p
		}
		n >>= 1
		m = ((m % p) * (m % p)) % p
	}
	return ans
}

// 快速加同理, 求m个n的值
func Quickadd(m, n int) int {
	ans := 0
	for m > 0 {
		if m&1 == 1 {
			ans = ans + n
		}
		m >>= 1
		n = n + n
	}
	return ans
}

// 如果有多组数据加入时就可以使用快速幂进行求解

// 方阵快速幂
// 两个矩阵相乘，n*m,m*p
func matmul(nums1, nums2 [][]int, n, m, p int) [][]int {
	// n := len(nums1)
	// m := len(nums2)
	// p := len(nums2[0])
	nums3 := make([][]int, n)
	for i := 0; i < len(nums3); i++ {
		nums3[i] = make([]int, p)
	}
	// 第一个矩阵的一行中的每个数去乘第二个矩阵的每一行
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nums1[i][j] != 0 {
				for k := 0; k < p; k++ {
					nums3[i][k] = nums1[i][j]*nums2[j][k] + nums3[i][k]
				}
			}
		}
	}
	return nums3
}

func matmulMod(nums1, nums2 [][]int, n, m, p int, h int) [][]int {
	// n := len(nums1)
	// m := len(nums2)
	// p := len(nums2[0])
	nums3 := make([][]int, n)
	for i := 0; i < len(nums3); i++ {
		nums3[i] = make([]int, p)
	}
	// 第一个矩阵的一行中的每个数去乘第二个矩阵的每一行
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nums1[i][j] != 0 {
				for k := 0; k < p; k++ {
					nums3[i][k] = ((nums1[i][j]*nums2[j][k])%h + (nums3[i][k])%h) % h
				}
			}
		}
	}
	return nums3
}

// p阶矩阵求n次幂
func matrixQuickPow(nums [][]int, p int, n int) [][]int {
	ans := make([][]int, p)
	for i := 0; i < p; i++ {
		ans[i] = make([]int, p)
		ans[i][i] = 1
	}
	for n > 0 {
		if n&1 == 1 {
			ans = matmul(ans, nums, p, p, p)
		}
		n >>= 1
		nums = matmul(nums, nums, p, p, p)
	}
	return ans
}

// everyday mat= [[0,1],[1,1]]
// [fn-1, fn] = mat的n-1 * [f0,f1]
// 与斐波那契数列有关

// 矩阵快速幂
func matQuickPowandMod(nums [][]int, p int, n int, k int) [][]int {
	ans := make([][]int, p)
	for i := 0; i < p; i++ {
		ans[i] = make([]int, p)
		ans[i][i] = 1
	}
	for n > 0 {
		if n&1 == 1 {
			ans = matmulMod(ans, nums, p, p, p, k)
		}
		n >>= 1
		nums = matmulMod(nums, nums, p, p, p, k)
	}
	return ans
}
