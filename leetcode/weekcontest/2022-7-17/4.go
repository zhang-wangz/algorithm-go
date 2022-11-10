package _022_7_17

import "sort"

func MinOperations(nums []int, numsDivide []int) int {
	maxd := numsDivide[0]
	for i := 1; i < len(numsDivide); i++ {
		num := gcd(maxd, numsDivide[i])
		if num == 1 {
			maxd = 1
			break
		} else {
			maxd = num
		}
	}
	m := map[int]bool{}
	sort.Ints(nums)

	if maxd < nums[0] {
		return -1
	}
	cnt := 0
	for i := 0; i < len(nums); i++ {
		if maxd%nums[i] == 0 {
			return cnt
		} else {
			if m[nums[i]] {
				cnt++
			} else {
				cnt++
				m[nums[i]] = true
			}
		}
	}
	return cnt
}

func gcd(a, b int) int {
	// 如果分子大于分母，需要交换顺序
	if a < b {
		a, b = b, a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
