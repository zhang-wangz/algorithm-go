package solution

// https://leetcode.cn/problems/height-checker/
// 学校打算为全体学生拍一张年度纪念照。根据要求，学生需要按照 非递减 的高度顺序排成一行。
//排序后的高度情况用整数数组 expected 表示，其中 expected[i] 是预计排在这一行中第 i 位的学生的高度（下标从 0 开始）。
//给你一个整数数组 heights ，表示 当前学生站位 的高度情况。heights[i] 是这一行中第 i 位学生的高度（下标从 0 开始）。
//返回满足 heights[i] != expected[i] 的 下标数量 。

func heightChecker(heights []int) int {
	ans := make([]int, len(heights))
	copy(ans, heights)
	cnt := 0
	quickSort(ans, 0, len(heights)-1)
	mergeSort(ans)
	for i := 0; i < len(heights); i++ {
		if ans[i] != heights[i] {
			cnt++
		}
	}
	return cnt
}

func quickSort(nums []int, start, end int) {
	if len(nums) <= 1 {
		return
	}
	if start < end {
		p := findP(nums, start, end)
		quickSort(nums, start, p-1)
		quickSort(nums, p+1, end)
	}
}

func findP(nums []int, start int, end int) int {
	p := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < p {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[end], nums[i] = nums[i], nums[end]
	return i
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums) / 2
	l := mergeSort(nums[:mid])
	r := mergeSort(nums[mid:])
	return merge(l, r)
}

func merge(left []int, right []int) (res []int) {
	l, r := 0, 0
	if left[l] < right[r] {
		res = append(res, left[l])
		l++
	} else {
		res = append(res, right[r])
		r++
	}
	res = append(res, left[l:]...)
	res = append(res, right[r:]...)
	return
}
