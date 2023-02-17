# 前缀和求解， zip(*g) 可以转置求列
from itertools import accumulate
from typing import List

# 这种题还是画图理解比较好
class Solution:
    def largest1BorderedSquare(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        rs = [list(accumulate(row, initial=0)) for row in grid]
        cs = [list(accumulate(col, initial=0)) for col in zip(*grid)]
        for d in range(min(m,n), 0, -1):
            for i in range(m-d+1):
                for j in range(n-d+1):
                    if rs[i][j+d] - rs[i][j] == d \
                        and rs[i+d-1][j+d] - rs[i+d-1][j] == d \
                        and cs[j][i+d] - cs[j][i] == d \
                        and cs[j+d-1][i+d] - cs[j+d-1][i] == d:
                        return d * d
        return 0