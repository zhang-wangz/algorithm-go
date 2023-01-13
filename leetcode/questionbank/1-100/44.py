# 通配符匹配与正则匹配那道题类似
# /?匹配单字符，*可以匹配任意字符
# dp

class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        m, n = len(s), len(p)
        # 初始化
        dp = [[False] * (n+1) for i in range(m+1)]
        dp[0][0] = True
        for i in range(1, n+1):
            if p[i-1] == "*":
                dp[0][i] = True
            else:
                break
        for i in range(1, m+1):
            for j in range(1, n+1):
                if p[j-1] == "*":  # 如果是*，就任意一个字段-1成立的话都成立
                    dp[i][j] = dp[i-1][j] | dp[i][j-1]
                elif p[j-1] == "?" or s[i-1] == p[j-1]:
                    dp[i][j] = dp[i-1][j-1]
        return dp[m][n]