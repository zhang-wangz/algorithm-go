class Solution:
    def maxValue(self, n: int, index: int, maxSum: int) -> int:
        def count(big, length):
            last = big - 1
            # 如果长度+1比最右边的还要小，就说明是三角形
            if length + 1 <= last:
                # 首+末 // 2
                small = last - (length-1)
                return (small + last) * length // 2
            else:
                # 等于1的正方形部分
                ones = length - last
                return  (1 + last) * last // 2 + ones

        l, r = 1, maxSum + 1
        while l < r:
            mid = (l + r) >> 1
            if mid + count(mid, index) + count(mid, n-1-index) <= maxSum:
                l = mid + 1
            else:
                r = mid
        return l-1
            
