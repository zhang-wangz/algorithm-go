from functools import cache


class Solution:
    def minDifficulty(self, a: List[int], d: int) -> int:
        n = len(a)
        if n < d:
            return -1
        # 定义f(i, j)为i天完成前j个任务的最小难度
        @cache
        def dfs(i, j):
            if i == 1:
                return max(a[:j+1])
            # if i == 0:
            #     return -1
            res, mx = inf, 0
            for k in range(j, i-2, -1):
                mx = max(mx, a[k])
                res = min(res, dfs(i-1, k-1)+mx)
            return res
        return dfs(d, n-1)