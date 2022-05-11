package solution

// 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：
// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。

// 我的思路是搜索行，然后搜索指定行，也可以把这个矩阵化成1维
// 即转化成 matrix[start/col][start%col]
// col = len(matrix[0])
func searchMatrix(matrix [][]int, target int) bool {
	var i int
	for i = 0; i < len(matrix); i++ {
		if target < matrix[i][0] {
			break
		}
	}
	if i == 0 {
		return false
	}
	start := i - 1
	nums := matrix[start]
	l := 0
	r := len(nums) - 1
	for j := 0; j < len(nums)-1; j++ {
		mid := l + (r-l)/2
		if nums[mid] < target {
			l = mid
		} else if nums[mid] > target {
			r = mid
		} else {
			return true
		}
	}
	if nums[l] == target {
		return true
	} else if nums[r] == target {
		return true
	}

	return false
}
