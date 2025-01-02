from itertools import accumulate
from typing import List
# 前缀和单调栈
class Solution:
    def longestWPI(self, hours: List[int]) -> int:
        n = len(hours)
        score = [1 if h > 8 else -1 for h in hours]
        presum = [0] + list(accumulate(score))
        ans = 0
        st = []
        # 生成单调栈
        for i in range(n+1):
            if not st or presum[st[-1]] > presum[i]:
                st.append(i)
        i = n 
        while i > ans:
            # 如果前缀此时大于里面的数，说明j-i大于0
            while st and presum[st[-1]] < presum[i]:
                ans = max(ans, i - st[-1])
                st.pop()
            i -= 1
        return ans

# if __name__ == '__main__':
#     s = Solution()
#     s.longestWPI([9,9,6,0,6,6,9])