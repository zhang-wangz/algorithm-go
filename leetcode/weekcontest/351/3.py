class Solution:
    def numberOfGoodSubarraySplits(self, nums: List[int]) -> int:
        mod = 1e9 + 7
        ans = 1
        cnt = 0
        d1 = False
        for n in nums:
            if n == 0:
                if d1:
                    cnt += 1
            else:
                d1 = True
                ans = ans * (cnt + 1) % mod
                cnt = 0
        return int(ans % mod) if d1  else 0