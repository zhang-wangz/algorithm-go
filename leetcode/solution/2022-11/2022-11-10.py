from collections import defaultdict, deque
from typing import List


class Solution:
    def shortestPathAllKeys(self, grid: List[str]) -> int:
        # 拿到所有的钥匙，求最短的距离
        dirs = [(1,0), (-1,0), (0,1), (0,-1)]
        dist = defaultdict(lambda: 0x3f3f3f3f) # 距离
        m, n, cnt = len(grid), len(grid[0]), 0
        q = deque([])
        for i, row in enumerate(grid):
            for j, c in enumerate(row):
                if c == '@':
                    # i, j, cur
                    dist[(i, j, 0)] = 0
                    # i, j, cur
                    q.append((i, j, 0))
                elif c.islower():
                    cnt += 1
        while q:
            x, y, cur = q.popleft()
            step = dist[(x,y,cur)]
            for dir in dirs:
                nx, ny = x + dir[0], y + dir[1]
                if nx < 0 or nx >= m  or ny < 0 or ny >= n:
                    continue
                c = grid[nx][ny]
                if c == '#':
                    continue
                if c.isupper() and cur >> ord(c)-ord('A') & 1 == 0:
                    continue
                ncur = cur
                if c.islower():
                    ncur |= 1 << ord(c) - ord('a')
                if ncur == (1 << cnt) -1: return step + 1
                if step + 1 >= dist[(nx, ny, ncur)]:
                    continue
                dist[(nx, ny, ncur)] = step + 1
                q.append((nx, ny, ncur))
            
        return -1