class Solution:
    def maxSumAfterPartitioning(self, arr: List[int], k: int) -> int:
        n = len(arr)
        dp = [0] * (n+1) # 前n个数组成的最大值(不含当前)
        for i in range(n+1):
            mx = 0
            j = i-1
            while i-j <= k and j>=0:
                mx = max(mx, arr[j])
                dp[i] = max(dp[i], dp[j]+(i-j)*mx)
                j -= 1
        return dp[-1]