package number

// https://leetcode.cn/problems/consecutive-numbers-sum/
// 给定一个正整数 n，返回 连续正整数满足所有数字之和为 n 的组数 。

// 等差数列求和公式
// sum = (a1+an)*len/2 = (a1+a1+(len-1)d)*len/2 = (2a1+len-1)*len/2
// sum * 2  = (2a1+len-1)*len
func consecutiveNumbersSum(n int) int {
	n = n << 1
	cnt := 0
	for l := 1; l <= n/2; l++ {
		if n%l == 0 && ((n/l-l+1)&1) == 0 {
			cnt++
		}
	}
	return cnt
}

// TLE
func consecutiveNumbersSum2(n int) int {
	sum := 0
	cnt := 0
	left, right := 0, 0
	for right <= n {
		right++
		sum += right
		if sum < n {
			continue
		}
		for sum >= n {
			if sum == n {
				cnt++
			}
			left++
			sum -= left
		}
	}
	return cnt
}
