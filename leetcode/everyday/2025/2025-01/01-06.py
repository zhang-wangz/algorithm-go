# -*- coding: utf-8 -*-
# @Time    : 2025/1/6 11:02 
# @File    : 01-06.py
from collections import defaultdict
from itertools import pairwise
from typing import List


# https://leetcode.cn/problems/maximum-consecutive-floors-without-special-floors/description/?envType=daily-question&envId=2025-01-06
# 算特殊楼层避免oom或tle
class Solution:
    def maxConsecutive(self, bottom: int, top: int, special: List[int]) -> int:
        special.sort()
        ans = max(special[0] - bottom, top - special[-1])
        for x, y in pairwise(special):
            ans = max(ans, y - x - 1)
        return ans