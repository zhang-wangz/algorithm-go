from typing import List
'''
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
'''
class Solution:
    def firstMissingPositive(self, nums: List[int]) -> int:
        # 所有<=0的数都取为n+1，取为正数
        n = len(nums)
        for i in range(n):
            if nums[i] <= 0:
                nums[i] = n + 1
        
        # 对每个数取abs，并且如果小于等于n的时候，将这个值所在的位置的值取负
        for i in range(n):
            num = abs(nums[i])
            if num <= n:
                nums[num - 1] = -abs(nums[num - 1])
        
        for i in range(n):
            if nums[i] > 0:
                return i + 1
        
        return n + 1