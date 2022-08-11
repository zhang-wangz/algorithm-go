package main

func findLength(nums1 []int, nums2 []int) int {
	n, m := len(nums1), len(nums2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if nums1[i] == nums2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = 0
			}
			ans = max(ans, dp[i+1][j+1])
		}
	}
	//for i := n - 1; i >= 0; i-- {
	//	for j := m - 1; j >= 0; j-- {
	//		if nums1[i] == nums2[j] {
	//			dp[i][j] = dp[i+1][j+1] + 1
	//		}else {
	//			dp[i][j] = 0
	//		}
	//		ans = max(ans, dp[i][j])
	//	}
	//}
	return ans
}

//func max(a, b int)int {
//	if a > b {
//		return a
//	}else {
//		return b
//	}
//}
