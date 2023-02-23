# 之前做过的一道题
from cgitb import grey
from typing import List

# 格雷码
class Solution:
    def circularPermutation2(self, n: int, start: int) -> List[int]:
        return [x ^ (x >> 1) ^ start for x in range(1<<n)]

    # 递归
    def circularPermutation(self, n: int, start: int) -> List[int]:
        def dfs(n):
            if n == 1:
                return ["0", "1"]
            grey = ["0"] * (1<<n)
            last = dfs(n-1)
            for i in range(len(last)):
                grey[i] = "0" + last[i]
                grey[len(grey)-i-1] = "1" + last[i]
            return grey 
        grey = dfs(n)
        for i in range(len(grey)):
            grey[i] = int(grey[i], 2) ^ start
        return grey
            
