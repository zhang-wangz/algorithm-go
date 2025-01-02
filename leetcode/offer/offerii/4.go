package main
// https://leetcode.cn/problems/search-a-2d-matrix-ii/
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for _, row := range matrix {
		l, r := 0, len(row)
		for l < r {
			mid := l + (r-l)>>1
			if row[mid] == target {
				return true
			} else if row[mid] > target {
				r = mid
			} else {
				l = mid + 1
			}
		}
	}
	return false
}
