package main

// 1191
// 给定一个整数数组arr和一个整数k，通过重复k次来修改数组。
// 例如，如果arr = [1, 2]，k = 3，那么修改后的数组将是 [1, 2, 1, 2, 1, 2] 。
// 返回修改后的数组中的最大的子数组之和。注意，子数组长度可以是 0，在这种情况下它的总和也是 0。
// 由于结果可能会很大，需要返回的109+ 7的模。
const mod = 1e9 + 7

func kConcatenationMaxSum(arr []int, k int) (ans int) {
	// kadane算法
	n := len(arr)
	sum := arr[0]
	maxEnd := 0
	if arr[0] > 0 {
		maxEnd = arr[0]
	}
	ans = maxEnd
	for i := 1; i < min(k, 2)*n; i++ {
		maxEnd = max(maxEnd+arr[i%n], arr[i%n])
		ans = max(maxEnd, ans)
		if i < n {
			sum += arr[i]
		}
	}
	for k = k - 1; k >= 2 && sum > 0; k-- {
		ans = (ans + sum) % mod
	}
	return ans % mod
}

//tle
func kConcatenationMaxSum2(arr []int, k int) (ans int) {
	if len(arr) == 0 {
		return 0
	}
	s := 0
	for i := 0; i < k; i++ {
		for _, v := range arr {
			s = s + v
			if s < 0 {
				s = 0
			}
			if s > ans {
				ans = s
				ans = ans % mod
			}
		}
	}
	return
}
