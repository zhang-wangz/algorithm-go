from typing import List


class Solution:
    def maxValue(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        dp = [[0] * (n+1) for _ in range(m+1)]
        for i, row in enumerate(grid):
            for j, x in enumerate(row):
                dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1]) + x
        return dp[m][n]