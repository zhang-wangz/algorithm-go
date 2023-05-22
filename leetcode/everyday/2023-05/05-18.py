class Solution:
    def addNegabinary(self, a1: List[int], a2: List[int]) -> List[int]:
        # 负进制的时候， 逢(2)进（-1），因为是负数, (-2)^i + (-2)^i = -1 * (-2)^(i+1)
        # 如果是-1的话，因为(-2)^i = (-2)^i + (-2)^(i+1)
        i, j = len(a1)-1, len(a2)-1
        c = 0
        ans = []
        while i >=0 or j >= 0 or c:
            a = 0 if i < 0 else a1[i]
            b = 0 if j < 0 else a2[j]
            x = a + b + c
            c = 0
            if x >= 2:
                x -= 2
                c -= 1
            elif x == -1:
                x = 1
                c += 1
            i, j = i-1, j-1
            ans.append(x)
        while len(ans) > 1 and ans[-1] == 0:
            ans.pop()
        return ans[::-1]