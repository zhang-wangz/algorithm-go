package _022_06

// 给你一个整数数组nums，将它重新排列成nums[0] < nums[1] > nums[2] < nums[3]...的顺序。
// 你可以假设所有输入数组都可以得到满足题目要求的结果。
// 快速选择，n复杂度，求一组元素的第K大值或者前K大值时

func WiggleSort(nums []int) {
	n := len(nums)
	mid := (n - 1) / 2
	quickSelect(nums, 0, n-1, mid)
	t := nums[mid]
	i, j := 0, 0
	k := n - 1
	for j < k {
		if nums[j] > t {
			for j < k && nums[k] > t {
				k--
			}
			nums[j], nums[k] = nums[k], nums[j]
			k--
		} else if nums[j] < t {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			j++
		}
	}
	left := append([]int{}, nums[:mid+1]...)
	right := append([]int{}, nums[mid+1:]...)

	for i := 0; i < len(left); i++ {
		nums[2*i] = left[len(left)-1-i]
	}
	for i := 0; i < len(right); i++ {
		nums[2*i+1] = right[len(right)-1-i]
	}
	return
}

func quickSelect(nums []int, start, end int, tidx int) {
	p := nums[end]
	i, j := start, start
	for j < end {
		if nums[j] < p {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j++
		} else {
			j++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	if i > tidx {
		quickSelect(nums, start, i-1, tidx)
	} else if i < tidx {
		quickSelect(nums, i+1, end, tidx)
	} else {
		return
	}
}
