# 排序 前缀和 二分搜索
class Solution:
    def answerQueries(self, nums: List[int], queries: List[int]) -> List[int]:
        f = list(accumulate(sorted(nums)))
        return [bisect_right(f, q) for q in queries]