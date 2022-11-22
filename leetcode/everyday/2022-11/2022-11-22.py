from math import lcm


class Solution:
    def nthMagicalNumber(self, n: int, a: int, b: int) -> int:
        # fx = x//a + x // b + x // c
        mod = 10**9+7
        c = lcm(a,b)
        l, r = min(a,b), n*min(a,b)
        while l < r:
            mid = (l+r) // 2
            cnt = mid // a + mid // b - mid // c # 递增
            if cnt >= n:
                r = mid 
            else:
                l = mid + 1
        return r % mod
        