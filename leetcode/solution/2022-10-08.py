from typing import List
# 给定两个大小相等的数组 nums1 和 nums2，nums1 
# 相对于 nums 的优势可以用满足 nums1[i] > nums2[i] 的索引 i 的数目来描述。
# 返回 nums1 的任意排列，使其相对于 nums2 的优势最大化。

class Solution:
    def advantageCount(self, nums1: List[int], nums2: List[int]) -> List[int]:
        n = len(nums1)
        nums1.sort()
        idxs = sorted(range(n), key=lambda x: nums2[x])
        ans =  [0] * n
        l, r = 0, n-1
        for x in nums1:
            if x > nums2[idxs[l]]:
                ans[idxs[l]] = x
                l += 1
            else:
                ans[idxs[r]] = x
                r -= 1
        return ans