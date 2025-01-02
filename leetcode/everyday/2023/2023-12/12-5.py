class Solution:
    # https://leetcode.cn/problems/minimum-fuel-cost-to-report-to-the-capital/description/
    # 2477. 到达首都的最少油耗
    def minimumFuelCost(self, roads: List[List[int]], seats: int) -> int:
        ans = 0
        g = defaultdict(list)
        for a, b in roads:
            g[a].append(b)
            g[b].append(a)
        def dfs(a:int, fa:int) -> int:
            nonlocal ans
            sz = 1
            for b in g[a]:
                if b != fa:
                    t = dfs(b, a)
                    ans += ceil(t/seats)
                    sz += t
            return sz
        dfs(0, -1)
        return ans