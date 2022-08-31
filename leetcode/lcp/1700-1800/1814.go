package main

// 统计数组中好数对的个数
// nums[i] + rev[j] = nums[j] + rev[i]
// nums[i] - rev[i] = nums[j] - rev[j]
func countNicePairs(nums []int) (ans int) {
	const mod = 1e9 + 7
	m := map[int]int{}
	for _, num := range nums {
		cnt := num - rev(num)
		ans += m[cnt]
		m[cnt]++
		ans %= mod
		m[cnt] %= mod
	}
	return
}
func rev(n int) (ans int) {
	for n != 0 {
		ans = ans*10 + n%10
		n /= 10
	}
	return
}
