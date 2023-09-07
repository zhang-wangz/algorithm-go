from itertools import accumulate


class seg():
    # go使用seg [] l,r,val 来代替了初始化（build的时候传入1，n)
    def __init__(self, n: int):
        self.n = n
        self.val = [0] * (n * 4)

    # 这里的l，r就是传入1，n，单点更新
    def update(self, o: int, l: int, r: int, idx: int, val: int) -> None:
        if l == r:
            self.val[o] += val
            return
        m = (l + r) // 2
        # 因为go build的时候传入l，r然后build左右部分是(l,m) | (m+1,r)
        if idx <= m: self.update(o<<1, l, m, idx, val)
        else: self.update(o<<1|1, m + 1, r, idx, val)
        self.val[o] = self.val[o<<1] + self.val[o<<1|1]

    def query(self, o: int, l: int, r: int, L: int, R: int) -> int:
        if L <= l and r <= R: return self.val[o]
        sum = 0
        m = (l + r) // 2
        if L <= m: sum += self.query(o<<1, l, m, L, R)
        if R > m: sum += self.query(o<<1|1, m + 1, r, L, R)
        return int(sum)

class Solution:
    # mle
    def countServers2(self, n: int, logs: List[List[int]], x: int, queries: List[int]) -> List[int]:
        # 线段树
        tree = seg(n+1)
        for sid, t in logs:
            tree.update(1, 1, n+1, t, 1)
        ans = []
        for q in queries:
            ans.append(n-tree.query(1, 1, n+1, q-x, q))
        return ans
        
    def countServers(self, n: int, logs: List[List[int]], x: int, queries: List[int]) -> List[int]:
        logs.sort(key=lambda x: x[1])
        tmax = logs[-1][1]
        h = [0] * tmax
        for sid, t in logs:
            h[t] += 1
        sc = list(accumulate(h, initial=1))
        ans = []
        for q in queries:
            l, r = q-5, q
            ans.append(h[r]-h[l])
        return ans







