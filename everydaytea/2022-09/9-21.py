import os, sys
from io import BytesIO, IOBase
from typing import Counter

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

input = lambda: sys.stdin.readline().strip()
ints = lambda: list(map(int, input().split()))
Int = lambda: int(input())

def queryInteractive(a, b, c):
    print('? {} {} {}'.format(a, b, c))
    sys.stdout.flush()
    return int(input())

def answerInteractive(x1, x2):
    print('! {} {}'.format(x1, x2))
    sys.stdout.flush()

inf = float('inf')
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
'''
输入 t (≤1e4) 表示 t 组数据，每组数据输入 n(≤2e5) 和一个长为 n 的数组 a (0≤a[i]≤n)。所有数据的 n 之和不超过 2e5。
每次操作，你可以把数组中的一个数加一。
定义 mex(a) 表示不在 a 中的最小非负整数。
定义 f(i) 表示使 mex(a) = i 的最小操作次数。如果无法做到，则 f(i) = -1。
输出 n+1 个数：f(0), f(1), ..., f(n)。
'''
t = int(input())
for _ in range(t):
    n = int(input())
    arr = list(map(int, input().split()))
    arr.sort()
    ret = [-1] * (n + 1)
    # 要使得a中的最小非负为i，需要a中存在>=0 <i的所有整数

    # f从0到n计算
    cnt = Counter(arr)
    st = []
    sum = 0
    for i in range(0, n+1):
        # 如果不存在比i小的值时候才需要去st拿
        if i > 0 and cnt[i-1] == 0:
            if len(st) == 0:
                break
            j = st.pop()
            sum += i-j-1
        ret[i] = sum + cnt[i] # 把当前i的都需要+1
        # 要把多余的数字加到st中
        while i > 0 and cnt[i-1] > 1:
            cnt[i-1] -= 1
            st.append(i-1)
    print(" ".join(map(str, ret)))

