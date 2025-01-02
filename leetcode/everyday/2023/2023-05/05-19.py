class Solution:
    # 计数dp， 还有一种是回溯
    def numTilePossibilities(self, tiles: str) -> int:
        # f i,j 为前i种字符构造长为j的序列有多少种
        count = Counter(tiles).values()
        n, m = len(tiles), len(count)
        f = [[0]* (n+1) for _ in range(m+1)]
        f[0][0] = 1
        for i, cnt in enumerate(count, 1):
            for j in range(n+1):
                for k in range(min(j,cnt)+1):
                    f[i][j] += f[i-1][j-k] * comb(j, k)
        return sum(f[m][1:])