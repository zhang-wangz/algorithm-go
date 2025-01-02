class Solution:
    # lis
    def makeArrayIncreasing(self, a: List[int], b: List[int]) -> int:
        b.sort()
        @cache
        def dfs(i, pre):
            if i < 0: return 0
            res = dfs(i-1, a[i]) if a[i] < pre else inf #不替换
            k = bisect_left(b, pre) - 1
            if k >= 0:
                res = min(res, dfs(i-1, b[k])+1)
            return res
        ans = dfs(len(a)-1, inf)
        return ans if ans < inf else -1
        