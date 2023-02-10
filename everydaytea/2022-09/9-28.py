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

# https://codeforces.com/problemset/problem/292/E

# 输入 n(≤1e5) 和 m (≤1e5)，两个长度都为 n 的数组 a 和 b（元素范围在 [-1e9,1e9] 内，下标从 1 开始）。
# 然后输入 m 个操作：
# 操作 1 形如 1 x y k，表示把 a 的区间 [x,x+k-1] 的元素拷贝到 b 的区间 [y,y+k-1] 上（输入保证下标不越界）。
# 操作 2 形如 2 x，输出 b[x]。
sys.setrecursionlimit(10**6)

# py超时，改成go应该可以过
def solve():
    n,m = ints()
    a = [0] + ints() # 都为n长度
    b = [0] + ints()

    v = [1e9] * (4*n)
    def query(o,l,r,i):
        if v[o] != 1e9 or l == r:
            return v[o]
        m = (l+r)//2
        if i <= m:
            return query(2*o,l,m,i)
        return query(2*o+1, m+1, r,i)

    def modify(o,l,r,x,y,val):
        if x<=l and r <=y:
            v[o] = val
            return
        # down
        if (d := v[o]) != 1e9:
            v[2*o] = d
            v[2*o+1] = d
            v[o] = 1e9
            
        m = (l+r)//2
        if x <= m:
            modify(2*o,l,m,x,y,val)
        if m < y:
            modify(2*o+1,m+1,r,x,y,val)
            
    for _ in range(m):
        op = ints()
        if op[0] == 1:
            x, y, k = op[1],op[2],op[3]
            modify(1,1,n,y,y+k-1,x-y)
        else:
            x = op[1]
            t = query(1,1,n,x)
            if t == 1e9:
                print(b[x])
            else:
                print(a[x+t])
solve()

