class Solution:
    def minHeightShelves(self, books: List[List[int]], shelfWidth: int) -> int:
        n = len(books)
        # 从0-i的书能够构成的最小高度
        @cache
        def dfs(i):
            if i < 0: return 0
            res, left_w, max_h = inf, shelfWidth, 0 
            for j in range(i, -1, -1):
                left_w -= books[j][0]
                if left_w < 0: break #放不下了就break，所有的结果已经在前面计算好了
                max_h = max(max_h, books[j][1]) #当前层放得下的话，继续放
                res = min(res, dfs(j-1)+max_h) #计算当前的最小
            return res
        return dfs(n-1)