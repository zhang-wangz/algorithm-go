# 评论区看到的解答, 最大的一定是最大的字母开始的，并且一定到最后，因为更长的一定排在更后面
class Solution:
    def lastSubstring2(self, s: str) -> str:
        c = max(s)
        ans = ''
        for i in range(len(s)):
            if s[i] == c:
                if s[i:] > ans:
                    ans = s[i:]
        return ans

    # 官解 同向双指针
    def lastSubstring(self, s: str) -> str:
        i, j, n = 0, 1, len(s)
        while j < n: # s[i:]是目前排在最后的，s[j:]是等待比较的
            k = 0 
            while j+k < n and s[i+k] == s[j+k]: # 前缀相等的s[i:]中一定有更长的
                k+=1
            if j+k < n and s[i+k] < s[j+k]:
                i, j = j, max(i+k+1, j+1)
            else:
                j = j+k+1 # 不相等就跳到下一个开头
        return s[i:]
