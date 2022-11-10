package solution

// 假设按照升序排序的数组在预先未知的某个点上进行了旋转
// ( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。 请找出其中最小的元素。
// [7,0,1,2,4,5,6]
func findMin(nums []int) int {
	// 把最后一个作为target
	if len(nums) == 0 {
		return -1
	}
	l := 0
	r := len(nums) - 1
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] <= nums[r] {
			r = mid
		} else {
			l = mid
		}
	}
	if nums[l] > nums[r] {
		return nums[r]
	}
	return nums[l]
}
