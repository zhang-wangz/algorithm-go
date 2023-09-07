from functools import cache


class Solution:
    def maximumANDSum(self, nums: List[int], numSlots: int) -> int:
        # mask, c的个数, 最大与和, 为0则被分配了
        n = numSlots * 2
        @cache
        def dfs(i, mask):
            if i >= len(nums):
                return 0
            ans = 0
            for j in range(n):
                if (mask >> j) & 1 == 0:
                    ans = max(ans, dfs(i+1, mask | (1<<j))+(nums[i]&(j//2+1)))
            return ans
        return dfs(0, 0)
    
    def maximumANDSum2(self, nums: List[int], numSlots: int) -> int:
        # 第i个数
        @cache
        def dfs(i, mask):
            if i < 0:
                return 0
            t, w, ans = mask, 1, 0
            for j in range(1, numSlots+1):
                if t % 3:
                    ans = max(ans, dfs(i-1, mask-w)+(nums[i]&j))
                t, w = t//3, w*3
            return ans
        return dfs(len(nums)-1, 3**numSlots-1)