class Solution:
    def restoreMatrix(self, rowSum: List[int], colSum: List[int]) -> List[List[int]]:
        m, n = len(rowSum), len(colSum)
        mat = [[0] * n for _ in range(m)]
        for i, rs in enumerate(rowSum):
            for j, cs in enumerate(colSum):
                mat[i][j] = x = min(rs, cs)
                rs -= x
                colSum[j] -= x
        return mat