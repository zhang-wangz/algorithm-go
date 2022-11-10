package solution

func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := MergeSort(nums[:mid])
	right := MergeSort(nums[mid:])
	res := merge(left, right)
	return res
}

func merge(l1, l2 []int) []int {
	m := len(l1)
	n := len(l2)
	res := make([]int, 0)
	var i, j int
	for i < m && j < n {
		if l1[i] < l2[j] {
			res = append(res, l1[i])
			i++
		} else {
			res = append(res, l2[j])
			j++
		}
	}
	res = append(res, l1[i:]...)
	res = append(res, l2[j:]...)
	return res
}
