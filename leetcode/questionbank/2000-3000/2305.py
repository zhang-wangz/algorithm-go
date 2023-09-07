class Solution:
    def distributeCookies(self, cookies: List[int], k: int) -> int:
        n = len(cookies)
        s = [0] * (1<<n)
        # 从小遍历,找到sum的值,相当于找到最后一个1去遍历状态
        for i in range(n):
            bit = 1<<i
            for j in range(bit):
                s[bit|j] = s[j]+cookies[i]
        
        dp = [[inf]*(1<<n) for _ in range(k)]
        dp[0] = s.copy()
        for i in range(1,k):
            for j in range(1, 1<<n):
                p = j
                while p:
                    tmp = max(s[p], dp[i-1][j^p])
                    dp[i][j] = min(dp[i][j], tmp)
                    p = (p-1)&j
        return dp[-1][-1]
