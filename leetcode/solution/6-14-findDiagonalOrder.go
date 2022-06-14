package solution

// https://leetcode.cn/problems/diagonal-traverse/
// 给你一个大小为 m x n 的矩阵 mat ，请以对角线遍历的顺序，用一个数组返回这个矩阵中的所有元素。

func findDiagonalOrder(mat [][]int) []int {
	res := make([]int, 0)
	if len(mat) <= 1 {
		return []int{mat[0][0]}
	}
	m := len(mat)
	n := len(mat[0])
	// i为偶数时，从下往上遍历
	// 如果i<m,则起点（i,0）
	// 如果i>=m,则起点(m-1,i-(m-1))

	// i为奇数时，从上往下遍历
	// 如果i<n,则起点(0,i)
	// 如果i>=n,则起点(i-(n-1),n-1)
	for i := 0; i < m+n-1; i++ {
		if i&1 == 0 {
			x0, y0 := 0, 0
			if i < m {
				x0 = i
				y0 = 0
			} else {
				x0 = m - 1
				y0 = i - (m - 1)
			}
			// 每次循环最大值为i，同时另一个值不能超过边界
			for x0 >= 0 && y0 < n {
				res = append(res, mat[x0][y0])
				x0--
				y0++
			}
		} else {
			x0, y0 := 0, 0
			if i < n {
				x0 = 0
				y0 = i
			} else {
				x0 = i - (n - 1)
				y0 = n - 1
			}
			for x0 < m && y0 >= 0 {
				res = append(res, mat[x0][y0])
				x0++
				y0--
			}
		}
	}
	return res
}
