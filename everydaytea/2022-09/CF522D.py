import os, sys
from io import BytesIO, IOBase
from types import GeneratorType

def bootstrap(f, stack=[]):
    def wrappedfunc(*args, **kwargs):
        if stack:
            return f(*args, **kwargs)
        else:
            to = f(*args, **kwargs)
            while True:
                if type(to) is GeneratorType:
                    stack.append(to)
                    to = next(to)
                else:
                    stack.pop()
                    if not stack:
                        break
                    to = stack[-1].send(to)
            return to
    return wrappedfunc

class TailRecurseException(BaseException):
    def __init__(self, args, kwargs):
        self.args = args
        self.kwargs = kwargs

def tail_call_optimized(g):
    """
    This function decorates a function with tail call
    optimization. It does this by throwing an exception
    if it is it's own grandparent, and catching such
    exceptions to fake the tail call optimization.
    This function fails if the decorated
    function recurses in a non-tail context.
    """
    def func(*args, **kwargs):
        f = sys._getframe()
        if f.f_back and f.f_back.f_back \
            and f.f_back.f_back.f_code == f.f_code:
            # 抛出异常
            raise TailRecurseException(args, kwargs)
        else:
            while 1:
                try:
                    return g(*args, **kwargs)
                except TailRecurseException as e:
                    args = e.args
                    kwargs = e.kwargs
    func.__doc__ = g.__doc__
    return func

class FastIO(IOBase):
    newlines = 0

    def __init__(self, file):
        self._fd = file.fileno()
        self.buffer = BytesIO()
        self.writable = "x" in file.mode or "r" not in file.mode
        self.write = self.buffer.write if self.writable else None

    def read(self):
        while True:
            b = os.read(self._fd, max(os.fstat(self._fd).st_size, 8192))
            if not b:
                break
            ptr = self.buffer.tell()
            self.buffer.seek(0, 2), self.buffer.write(b), self.buffer.seek(ptr)
        self.newlines = 0
        return self.buffer.read()

    def readline(self):
        while self.newlines == 0:
            b = os.read(self._fd, max(os.fstat(self._fd).st_size, 8192))
            self.newlines = b.count(b"\n") + (not b)
            ptr = self.buffer.tell()
            self.buffer.seek(0, 2), self.buffer.write(b), self.buffer.seek(ptr)
        self.newlines -= 1
        return self.buffer.readline()

    def flush(self):
        if self.writable:
            os.write(self._fd, self.buffer.getvalue())
            self.buffer.truncate(0), self.buffer.seek(0)

class IOWrapper(IOBase):
    def __init__(self, file):
        self.buffer = FastIO(file)
        self.flush = self.buffer.flush
        self.writable = self.buffer.writable
        self.write = lambda s: self.buffer.write(s.encode("ascii"))
        self.read = lambda: self.buffer.read().decode("ascii")
        self.readline = lambda: self.buffer.readline().decode("ascii")

sys.stdin, sys.stdout = IOWrapper(sys.stdin), IOWrapper(sys.stdout)
if sys.hexversion == 50858480:
    sys.stdin = open('/Users/zhang-wangz/Desktop/algorithm/tea_0x3f/00_cfinput.txt')

def queryInteractive(a, b, c):
    print('? {} {} {}'.format(a, b, c))
    sys.stdout.flush()
    return int(input())

def answerInteractive(x1, x2):
    print('! {} {}'.format(x1, x2))
    sys.stdout.flush()

input = lambda: sys.stdin.readline().strip()
ints = lambda: list(map(int, input().split()))
Int = lambda: int(input())
inf = float('inf')

'''
输入 n(2≤n≤2e5) 和 m(1≤m≤2e5)；然后输入一个长为 n 的数组 a(-1e9≤a[i]≤1e9)，数组下标从 1 开始；
然后输入 m 个询问，每个询问表示一个数组 a 内的闭区间 [L,R] (1≤L≤R≤n)。
对每个询问，输出区间内的相同元素下标之间的最小差值，如果区间内不存在相同元素，输出 -1。
'''
# tle 3000ms
class ST:
    # go使用seg [] l,r,val 来代替了初始化（build的时候传入1，n)
    def __init__(self, n: int):
        self.n = n
        self.val = [10**9] * (n * 4)

    # 这里的l，r就是传入1，n，这里是单值更新
    # @tail_call_optimized
    def update(self, o: int, l: int, r: int, idx: int, val: int) -> None:
        if l == r:
            self.val[o] = val
            return
        m = (l + r) // 2
        # 因为go build的时候传入l，r然后build左右部分是(l,m) | (m+1,r)
        if idx <= m: self.update(o<<1, l, m, idx, val)
        else: self.update(o<<1|1, m + 1, r, idx, val)
        self.val[o] = min(self.val[o<<1], self.val[o<<1|1])
    
    # @tail_call_optimized
    def query(self, o: int, l: int, r: int, x: int, y: int) -> int:
        if x <= l and r <= y: return self.val[o]
        m = (l + r) // 2
        # R <= m 说明最小值在左边
        if y <= m: return self.query(o<<1, l, m, x, y)
        # m < L 说明最小值在右边
        if m < x: return self.query(o<<1|1, m + 1, r, x, y)
        return min(self.query(o<<1, l, m, x, y), self.query(o<<1|1, m + 1, r, x, y))

# 1684ms 树状数组
class BinIndexTreeMin:
    def __init__(self, size):
        self.size = size
        self.a = [10**9 for _ in range(size+5)]
        self.h = self.a[:]
        self.mn = 10**9

    def update(self, x, v):
        if v < self.mn:
            self.mn = v
        a = self.a
        h = self.h
        a[x] = v
        while x <= self.size:
            if h[x] >= v:
                h[x] = v
            else:
                break
            x += self.lowbit(x)

    def query(self, l, r):
        ans = self.a[r]
        while l != r:
            r -= 1
            while r - self.lowbit(r) > l:
                if ans > self.h[r]:
                    ans = self.h[r]
                    if ans == self.mn:
                        break
                # ans = min(ans, self.h[r])
                r -= self.lowbit(r)
            # ans = min(ans, self.a[r])
            if ans > self.a[r]:
                ans = self.a[r]
            if ans == self.mn:
                break
        return ans

    def lowbit(self, x):
        return x & -x

def solve():
    n, m = ints()
    a = [0] + ints() # 从1开始
    ret = [-1] * m
    # q是记录右端点的
    q = [[] for _ in range(n+1)]
    for i in range(m):
        l, r = ints()
        # 每部分r记录l和此时的i，以右端点为序
        q[r].append((l, i))
    
    # sg = ST(n)
    sg = BinIndexTreeMin(n)
    last = {}
    # 1-n 之间进行遍历
    for i in range(1, n+1):
        tp = last.get(a[i])
        if tp and tp > 0:
            # sg.update(1, 1, n, tp, i-tp)
            # idx，val，更新最小值，这里的l,r是1到n，所以这是给下标为tp的那一段设置为i-tp
            sg.update(tp, i-tp)
        for p in q[i]:
            if p[0] > i: break
            # res = sg.query(1, 1, n, p[0], i)
            # p[0],查询l到i的最小值，如果是inf，则没有该值
            res = sg.query(p[0], i)
            if res == 10**9:
                res = -1
            ret[p[1]] = res
        # 上次同元素的值序列是i
        last[a[i]] = i
    print("\n".join(map(str, ret)))
solve()