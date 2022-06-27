package offer

// NumMatrix https://leetcode.cn/problems/O4NDxx/
type NumMatrix struct {
	sum [][]int
}

func Constructor1(matrix [][]int) NumMatrix {
	ans := make([][]int, len(matrix))
	for i := 0; i < len(ans); i++ {
		ans[i] = make([]int, len(matrix[i])+1)
		ans[i][0] = 0
		for j := 1; j <= len(matrix[i]); j++ {
			ans[i][j] = ans[i][j-1] + matrix[i][j-1]
		}
	}
	return NumMatrix{
		sum: ans,
	}
}

func (mat *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for j := row2; j >= row1; j-- {
		sum += mat.sum[j][col2+1] - mat.sum[j][col1]
	}
	return sum
}
