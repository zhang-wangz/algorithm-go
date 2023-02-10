class Solution:
    #dp dp[i][p][k]. i次 第i次是啥 重复了几次
    def dieSimulator(self, n: int, rollMax: List[int]) -> int:
        mod = int(1e9+7)
        rollMax = [0] + rollMax
        dp = [[[0] * 16 for _ in range(7)] for _ in range(n+5)]
        for i in range(1,7):
            dp[1][i][1] = 1
        for i in range(2, n+1):
            for j in range(1,7): # 上一次是啥
                for k in range(1, rollMax[j]+1):
                    for p in range(1,7): # 这一次是啥
                        if p != j:
                            dp[i][p][1] = (dp[i][p][1] + dp[i-1][j][k]) % mod
                        elif k + 1 <= rollMax[p]:
                            dp[i][p][k+1] = (dp[i][p][k+1] + dp[i-1][p][k]) % mod
        res = 0
        for i in range(1,7):
            for j in range(1, rollMax[i]+1):
                res = (res + dp[n][i][j]) % mod
        return res
                            
        



    # dfs
    def dieSimulator2(self, n: int, rollMax: List[int]) -> int:
        mod, rollMax = int(1e9+7), [0] + rollMax
        @cache
        def dfs(i, last, cnt):
            if i == n: return 1
            res = 0
            for k in range(1,7):
                if k == last:
                    if (cnt + 1) <= rollMax[k]:
                        res = (res%mod + dfs(i+1,k,cnt+1)%mod) % mod
                else:
                    res = (res%mod + dfs(i+1,k,1)%mod) % mod
            return res
        return dfs(0, 0, 0)

