package _022_7_17

import "sort"

func maximumSum(nums []int) int {
	m := map[int][]int{}
	for i := range nums {
		if m[compute(nums[i])] == nil {
			m[compute(nums[i])] = make([]int, 0)
		}
		m[compute(nums[i])] = append(m[compute(nums[i])], nums[i])
	}
	max := -1
	for k, _ := range m {
		sort.Slice(m[k], func(i, j int) bool {
			return m[k][i] > m[k][j]
		})
		if len(m[k]) >= 2 && m[k][0]+m[k][1] > max {
			max = m[k][0] + m[k][1]
		}
	}
	return max
}

func compute(num int) int {
	ans := 0
	for num != 0 {
		n := num % 10
		ans += n
		num /= 10
	}
	return ans
}
