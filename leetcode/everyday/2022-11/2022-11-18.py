from typing import List


class Solution:
    def sumSubseqWidths(self, nums: List[int]) -> int:
        mod = 10 ** 9 + 7
        n = len(nums)
        nums.sort()
        pow2 = [0] * n
        pow2[0] = 1
        for i in range(1,n):
            pow2[i] = pow2[i-1] * 2 % mod # 预处理2的幂次
        return sum((pow2[i] - pow2[-1-i]) * x 
                for i, x in enumerate(nums)) % mod 