package main

func lenLongestFibSubseq(arr []int) (ans int) {
	n := len(arr)
	dp := make([][]int, n)
	for f := range dp {
		dp[f] = make([]int, n)
	}
	for i := 2; i < n; i++ {
		j, k := 0, i-1
		for j < k {
			if arr[j]+arr[k] == arr[i] {
				if dp[k][i] == 0 {
					dp[k][i] = 3
				} else {
					dp[k][i] = max(dp[k][i], dp[j][k]+1)
				}
				ans = max(ans, dp[k][i])
				j++
				k--
			} else if arr[j]+arr[k] < arr[i] {
				j++
			} else {
				k--
			}
		}
	}
	return
}
