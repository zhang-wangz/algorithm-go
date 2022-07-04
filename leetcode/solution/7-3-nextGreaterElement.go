package solution

import (
	"math"
	"sort"
	"strconv"
)

// https://leetcode.cn/problems/next-greater-element-iii/
// 给你一个正整数n ，请你找出符合条件的最小整数，其由重新排列 n 中存在的每位数字组成，并且其值大于 n 。如果不存在这样的正整数，则返回 -1 。
//
//注意 ，返回的整数应当是一个 32 位整数 ，如果存在满足题意的答案，但不是 32 位整数 ，同样返回 -1 。
//

// NextGreaterElement 官方
func NextGreaterElement(n int) int {
	nums := []byte(strconv.Itoa(n))
	if len(nums) == 0 || len(nums) == 1 {
		return -1
	}
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i < 0 {
		return -1
	}
	// 找一个大于的最小数
	j := len(nums) - 1
	for j >= 0 && nums[j] <= nums[i] {
		j--
	}
	nums[j], nums[i] = nums[i], nums[j]
	reverse(nums[i+1:])
	ans, _ := strconv.Atoi(string(nums))
	if ans > math.MaxInt32 {
		return -1
	}
	return ans
}
func reverse(nums []byte) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}

func NextGreaterElement2(n int) int {
	MAX := math.MaxInt32
	v := strconv.Itoa(n)
	if len(v) == 0 || len(v) == 1 {
		return -1
	}
	bytev := []byte(v)
	f := false
	j := len(bytev) - 2
	for ; j >= 0; j-- {
		if bytev[j] < bytev[j+1] {
			f = true
			break
		}
	}
	// 从j+1，len中找出一个比j大的最小值
	if f {
		ml := math.MaxInt32
		idx := -1
		for i := len(bytev) - 1; i >= j+1; i-- {
			if bytev[i] > bytev[j] && int(bytev[i]-'0') < ml {
				ml = int(bytev[i] - '0')
				idx = i
			}
		}
		bytev[idx], bytev[j] = bytev[j], bytev[idx]
		nx := bytev[j+1:]
		sort.Slice(nx, func(i, h int) bool {
			return nx[i] < nx[h]
		})
		copy(bytev[j+1:], nx)
		n, _ := strconv.Atoi(string(bytev))
		if n <= MAX {
			return n
		} else {
			return -1
		}
	}
	return -1
}
