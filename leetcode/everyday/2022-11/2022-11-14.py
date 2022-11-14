from typing import List

## 左右侧分别取一半的值进行判断
## 通过*n 来消除浮点数
class Solution:
    def splitArraySameAverage(self, nums: List[int]) -> bool:
        n = len(nums)
        if n == 1:
            return False
        m = n // 2
        s = sum(nums)
        for i in range(n):
            nums[i] = nums[i] * n - s
        
        left = set()
        for i in range(1, 1<<m):
            tot = sum(x for j,x in enumerate(nums[:m]) if i >> j & 1)
            if tot == 0:
                return True
            left.add(tot)
        
        rsum = sum(nums[m:])
        for i in range(1, 1<<(n-m)):
            tot = sum(x for j,x in enumerate(nums[m:]) if i >> j & 1)
            if tot == 0 or rsum != tot and  -tot in left:
                return True
        return False