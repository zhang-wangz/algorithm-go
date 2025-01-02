# https://leetcode.cn/problems/parallel-courses-ii/solutions/2310832/python3javacgo-yi-ti-yi-jie-zhuang-tai-y-dabh/
from functools import cache
from itertools import combinations
from typing import List


class Solution:
    def minNumberOfSemesters(self, n: int, relations: List[List[int]], k: int) -> int:
        
        @cache
        def dfs(mask):
            if mask == (1<<n)-1:
                return 0
            nex = [i for i in range(n) if not mask & (1<<i) and mask & pre[i] == pre[i]]
            res = n + 1
            for sub in combinations(nex, min(k, len(nex))):
                res = min(res, 1 + dfs(mask | sum([1<<i for i in sub])))
            return res
        pre = [0] * n
        for i, j in relations:
            pre[j-1] |= 1 << (i-1)
        return dfs(0)
        

