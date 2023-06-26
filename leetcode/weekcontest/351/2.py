from itertools import count


class Solution:
    def makeTheIntegerZero(self, num1: int, num2: int) -> int:
        # 从暴力入手:枚举操作次数 k
        # 问题变成:设 x = num1 - num2 * k
        # 计算 x 能否分解成 k 个 2^i
        # k in [x.bit count()，x]
        for k in count(1):
            x = num1 - num2 * k
            if x < k: return -1
            if k >= x.bit_count(): return k
        