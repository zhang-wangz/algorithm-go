from xml.dom.minidom import Node


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
                if cur.child[bit]:
                    res += cur.child[bit].cnt
                cur = cur.child[bit ^ 1]
            else:
                cur = cur.child[bit]
        return res



class Solution:
    def countPairs(self, nums: List[int], low: int, high: int) -> int:
        ans = 0
        t = trieXor()
        for num in nums:
            ans += t.find(num, high + 1) - t.find(num, low)
            t.insert(num)
        return ans