from math import comb

class Solution:
    def countVowelStrings2(self, n: int) -> int:
        # 关于我高中的知识，做力扣才搞懂这件事
        # n个球放到5个盒子里有几种放法，可以允许部分盒子里是空的，不放球
        # 挡板在两边是不能放的，不然盒子都凑在一起，相当于没分配，所以空位要-2
        # 相当于 m+n+1-2 = 5+n+1-2 = n+4, 选4个位置, 先借后还的思想
        return comb(n+4, 4)

    def countVowelStrings(self, n: int) -> int:
        dp = [1] * 5
        for _ in range(n-1): # 已经分配了一个球，所以只需要分配n-1个
            for j in range(1,5): # 同理，以'0'对应的'a'结尾只可能有一个，不需要管，所以从'b'结尾开始计算
                dp[j] += dp[j-1]
        return sum(dp)