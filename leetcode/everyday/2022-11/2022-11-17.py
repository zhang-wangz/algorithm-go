# 子序列个数
# 二分或者字符串匹配
from bisect import bisect_right
from collections import defaultdict
from typing import List


class Solution:
    def numMatchingSubseq(self, s: str, words: List[str]) -> int:
        # 1.二分
        dic = defaultdict(list)
        for i,c in enumerate(s):
            dic[c].append(i)
        ans = len(words)
        # 依次判断每个字符串
        for w in words:
            if len(w) > len(s):
                ans -= 1
                continue
            p = -1
            for c in w:
                ps = dic[c]
                j = bisect_right(ps, p)
                if j == len(ps):
                    ans -= 1
                    break
                p = ps[j]
                
        # 2.指针匹配
        dic = defaultdict(list)
        for i, w in enumerate(words):
            dic[w[0]].append((i, 0))
        ans = 0
        for c in s:
            cs = dic[c]
            dic[c] = []
            for i,j in cs:
                j += 1
                if j == len(words[i]):
                    ans += 1
                    continue
                dic[words[i][j]].append((i, j))
        
        return ans
            