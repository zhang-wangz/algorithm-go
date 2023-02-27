class Solution:
    def movesToMakeZigzag(self, nums: List[int]) -> int:
        s = [0, 0]
        for i in range(len(nums)):
            left = nums[i-1] if i else inf
            right = nums[i+1] if i < len(nums)-1 else inf
            s[i % 2] += max(nums[i] - min(left, right) + 1, 0)
        return min(s)