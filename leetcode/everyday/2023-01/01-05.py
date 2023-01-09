

from collections import Counter
from typing import List


BitLen = 14
class trieNode:
    def __init__(self) -> None:
        self.child = [None, None]
        self.cnt = 0

class trieXor:
    def __init__(self) -> None:
        self.root = trieNode()
    
    def insert(self, num: int):
        cur = self.root
        for k in range(BitLen, -1, -1):
            x = (num >> k) & 1
            if not cur.child[x]:
                cur.child[x] = trieNode()
            cur = cur.child[x]
            cur.cnt += 1
        
    def find(self, num, x) -> int:
        res = 0
        cur = self.root
        for k in range(BitLen, -1, -1):
            bit = (num >> k) & 1
            if cur == None: return res

            if x >> k & 1 :
                # 选择路径上与ai 第k位相等的子节点，xor之后是0，小于1
                if cur.child[bit]:
                    res += cur.child[bit].cnt
                # 为了路径上的值 xor之后与x的前缀保持一致
                cur = cur.child[bit ^ 1]
            else:
                cur = cur.child[bit ^ 0]
        return res


class Solution:
    def countPairs(self, nums: List[int], low: int, high: int) -> int:
        ans = 0
        t = trieXor()
        for num in nums:
            ans += t.find(num, high + 1) - t.find(num, low)
            t.insert(num)
        return ans

    def countPairs2(self, nums: List[int], low: int, high: int) -> int:
        ans, cnt = 0, Counter(nums)
        high += 1
        while high:
            nxt = Counter()
            for x, c in cnt.items():
                if high & 1: ans += c * cnt[x ^ (high-1)]
                if low & 1:  ans -= c * cnt[x ^ (low-1)]
                nxt[x >> 1] += c
            cnt = nxt
            low >>= 1
            high >>= 1
        return ans // 2

