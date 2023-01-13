# hard 37 解数独 回溯
from typing import List


class Solution:
    def solveSudoku(self, board: List[List[str]]) -> None:
        def dfs(pos: int):
            nonlocal valid
            if pos == len(space):
                valid = True
                return
            i, j = space[pos]
            for num in range(9):
                if line[i][num] == column[j][num] == block[i//3][j//3][num] == False:
                    line[i][num] = column[j][num] = block[i//3][j//3][num] = True
                    board[i][j] = str(num + 1)
                    dfs(pos + 1)
                    line[i][num] = column[j][num] = block[i//3][j//3][num] = False
                if valid:
                    return

        line = [[False] * 9 for _ in range(9)]
        column = [[False] * 9 for _ in range(9)]
        block =[[[False]* 9 for _a in range(3)] for _b in range(3)]
        space = []
        valid = False
    
        for i in range(9):
            for j in range(9):
                if board[i][j] == '.':
                    space.append((i,j))
                else:
                    num = int(board[i][j])-1
                    line[i][num] = column[j][num] = block[i//3][j//3][num] = True
        dfs(0)
                    
                
