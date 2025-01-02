# 1590. 使数组和能被 P 整除
from itertools import accumulate
from typing import List


class Solution:
    def minSubarray(self, nums: List[int], p: int) -> int:
        s = list(accumulate(nums, initial=0))
        x = s[-1] % p  # 数组和
        if x == 0: return 0
        dis = {}
        ans = n = len(nums)
        for i, v in enumerate(s):
            dis[v % p] = i
            jv = (v-x) % p
            if jv in dis:
                ans = min(ans, i-dis[jv])
        return ans if ans < n else -1    
            
            