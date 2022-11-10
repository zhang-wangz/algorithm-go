package sort

func MergeSort(nums []int) []int {
	return mergeSort(nums)
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])
	res := merge(left, right)
	return res
}

func merge(left, right []int) (res []int) {
	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			res = append(res, right[r])
			r++
		} else {
			res = append(res, left[l])
			l++
		}
	}
	res = append(res, left[l:]...)
	res = append(res, right[r:]...)
	return res
}
