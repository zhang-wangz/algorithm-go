class Solution:
    # n2 摆烂
    def maximalNetworkRank(self, n: int, roads: List[List[int]]) -> int:
        grp = [[False] * n for _ in range(n)]
        degree = [0] * n
        for a,b in roads:
            degree[a] += 1
            degree[b] += 1
            grp[a][b] = True
            grp[b][a] = True
        ans = 0
        for i in range(n):
            for j in range(i+1, n):
                rank = degree[i] + degree[j] - grp[i][j]
                ans = max(ans, rank)
        return ans