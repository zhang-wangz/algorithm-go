from typing import List


# 计数
class Solution:
    def numSubarrayBoundedMax(self, nums: List[int], left: int, right: int) -> int:
        def count(lower: int) -> int:
            res = cur = 0
            for x in nums:
                if x <= lower:
                    # 每多一个数组就数字+1
                    cur += 1
                else:
                    cur = 0
                res += cur
            return res
        return count(right) - count(left-1)
        