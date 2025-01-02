class Solution:
    def mostFrequentEven(self, nums: List[int]) -> int:
        cnt = Counter(nums)
        ans = inf
        for k, v in cnt.items():
            if k & 1 == 0 and cnt[k] >= cnt[ans]:
                if cnt[k] == cnt[ans]:
                    if k < ans:
                        ans = k
                else:
                    ans = k    
        return -1 if ans == inf else ans