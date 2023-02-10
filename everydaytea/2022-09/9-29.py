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

# https://codeforces.com/problemset/problem/558/C

# 输入 n(≤1e5) 和一个长为 n 的数组 a (1≤a[i]≤1e5)。
# 每次操作你可以将某个 a[i] 乘二或除二（下取整）。
# 输出至少需要多少次操作，可以让 a 的所有数都相同。
def solve2():
    n = Int()
    a = ints()
    m = max(a) + 1
    # 记录x已经存在多少个了
    cnt = [0] * m
    # 所有数变成x花了多少步
    steps = [0] * m
    # 是否和上次x时的i一致，作为退出条件
    vis = [-1] * m
    # 遍历
    for i, v in enumerate(a):
        q = [(v, 0)]
        step = 0
        while step < len(q):
            v, s = q[step]
            step += 1
            if v >= m or vis[v] == i:
                continue
            vis[v] = i  #防止重复
            steps[v] += s  #把步数加起来
            cnt[v] += 1  #出现次数加起来
            q.append((v * 2, s + 1))
            q.append((v // 2, s + 1))
    ans = min(steps[x] for x in range(m) if cnt[x] == n) #只统计等于n的
    print(ans)

def solve():
    n = Int()
    a = ints()
    m = max(a) + 1
    sum = [0] * m
    cnt = [0] * m
    for _, v in enumerate(a):
        c = 0
        while v:
            cnt[v] += 1
            sum[v] += c
            c += 1
            v //= 2
    v = m-1
    while cnt[v] < n:
        v -= 1
    
    ans = sum[v]
    t = v * 2
    # 如果左移一位减少的操作数大于此时n-cnt[t]增加的操作数，就减去这些操作数
    while t < m and cnt[t] > n - cnt[t]:
        ans -= 2 * cnt[t] - n
        t = t * 2
    print(ans)



if __name__ == '__main__':
    solve()