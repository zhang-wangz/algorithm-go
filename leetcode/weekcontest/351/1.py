from math import gcd

class Solution:
    def countBeautifulPairs2(self, nums: List[int]) -> int:
        ans = 0
        n = len(nums)
        for i in range(n):
            x = nums[i]
            while x >= 10: x //= 10
            for j in range(i+1, n):
                last = nums[j] % 10
                if gcd(last, x) == 1:
                    ans += 1
        return ans
    
    def countBeautifulPairs(self, nums: List[int]) -> int:
        ans = 0
        n = len(nums)
        cnt = [0] * n
        for x in nums:
            for j in range(1, 10):
                if gcd(x % 10, j) == 1:
                    ans += cnt[j]
            while x >= 10: x //= 10
            cnt[x] += 1
        return ans  
        
