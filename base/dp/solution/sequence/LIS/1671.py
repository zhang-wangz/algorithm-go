# https://leetcode.cn/problems/minimum-number-of-removals-to-make-mountain-array/description/

class Solution:
    def minimumMountainRemovals(self, nums: List[int]) -> int:
        # 从左到右看最长递增序列
        # 求递减子序列
        g = []
        n = len(nums)
        l, r = [0]*n, [0]*n
        for i,x in enumerate(nums):
            j = bisect_left(g, x)
            if j == len(g):
                g.append(x)
            else:
                g[j] = x
            l[i] = j+1
        g = []
        for i,x in enumerate(nums[::-1]):
            j = bisect_left(g,x)
            if j == len(g):
                g.append(x)
            else:
                g[j] = x
            r[n-1-i] = j+1
        sidx, send = 0, n-1
        for i in range(n):
            if i != 0 and nums[i] > nums[i-1]:
                sidx = i
                break
        for j in range(n-1, -1, -1):
            if j != n-1 and nums[j] > nums[j+1]:
                send = j
                break
        print(sidx,send)
        a = [i+1-l[i] + (n-i)-r[i] for i in range(sidx,send+1)]
        if len(a) == 0:
            return 0
        return min(a)
        