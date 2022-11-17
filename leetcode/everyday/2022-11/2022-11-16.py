from typing import List


class Solution:
    def isIdealPermutation(self, nums: List[int]) -> bool:
        # 找是否存在非局部倒置即可
        n = len(nums)
        m_suffix = nums[-1]
        for i in range(n-2, 0, -1):
            if nums[i-1] > m_suffix:
                return False
            m_suffix = min(m_suffix, nums[i])
        return True