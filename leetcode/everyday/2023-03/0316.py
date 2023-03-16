from typing import Counter, List

# 子数组 等价替换原则 中位数 (小于的数=大于的数， 左小+右小=右大+右大，左小-左大=右大-右小)
# 偶数的时候统计 右大-右小-1
class Solution:
    def countSubarrays(self, nums: List[int], k: int) -> int:
        index = nums.index(k)
        cnt, x = Counter({0:1}), 0
        for i in range(index-1, -1, -1):
            x += 1 if nums[i] < k else -1
            cnt[x] += 1
        
        ans, x = cnt[0] + cnt[-1], 0
        for i in range(index+1, len(nums)):
            x += 1 if nums[i] > k else -1
            ans += cnt[x] + cnt[x-1]
        return ans