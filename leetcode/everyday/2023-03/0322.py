from typing import List


class Solution:
    def bestTeamScore(self, scores: List[int], ages: List[int]) -> int:
        a = sorted(zip(scores, ages))
        f = [0] * len(a)
        for i in range(len(a)):
            for j in range(i):
                if a[j][1] <= a[i][1]:
                    f[i] = max(f[j], f[i])  # 如果分数<=的话， age也<=
            f[i] += a[i][0]
        return max(f)
                
    