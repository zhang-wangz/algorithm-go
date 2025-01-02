class Solution:
    def baseNeg2(self, n: int) -> str:
        if n == 0 or n == 1:
            return str(n)
        k = -2
        ans = []
        while n != 0:
            d = n % abs(k)
            ans.append(str(d))
            n -= d
            n //= -2
        return "".join(ans[::-1])