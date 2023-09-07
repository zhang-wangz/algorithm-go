
from heapq import heappop


class Solution:
    def findValueOfPartition(self, nums: List[int]) -> int:
        ans = inf
        for a,b in pairwise(sorted(nums, reverse=True)):
            ans = min(ans, a-b)
        return ans 