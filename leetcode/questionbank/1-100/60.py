class Solution:
    # 康托展开和逆康托展开
    # 求在对应的n个数组成的排列中，小于当前排列的数字有几个
    # https://blog.csdn.net/wbin233/article/details/72998375
    # 逆康托展开
    def getPermutation(self, n: int, k: int) -> str:
        # 求阶乘值
        factorial = [1]
        for i in range(1, n):
            factorial.append(factorial[-1] * i)
        
        # -1之后才是逆托康展开的值
        k -= 1
        ans = []
        valid = [1] * (n + 1)
        for i in range(1, n + 1): # [1,n]
            order = k // factorial[n - i] + 1 
            for j in range(1, n + 1):
                order -= valid[j] # 选择商+1的值，因为后面的那个才是需要的那个值，小于其的值为商的值
                if order == 0:
                    ans.append(str(j)) 
                    valid[j] = 0
                    break
            # 求下一个值的和        
            k %= factorial[n - i]

        return "".join(ans)