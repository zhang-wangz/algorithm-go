
from heapq import heappop, heappush
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
    sys.stdin = open('/Users/zhang-wangz/Desktop/algorithm/0x3f/cfinput.txt')

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
输入 n(2≤n≤2e5) 和 m(1≤m≤2e5)，然后输入 n 行，每行有两个数，表示一个闭区间（1≤L≤R≤1e18），
然后输入一个长为 m 的数组 a (1≤a[i]≤1e18)。
输入保证区间之间没有交集，且上一个区间的右端点小于下一个区间的左端点。

你有 m 座桥，每座桥的长为 a[i]，你需要选择 n-1 座桥连接所有相邻的区间（注意是相邻）。
要求桥的两个端点分别落在这两个闭区间内（这两个端点的差等于桥长）。

如果无法做到，输出 No；否则输出 Yes，然后按顺序输出这 n-1 座桥的编号（编号从 1 开始），
输出的第一座桥连接第一个区间和第二个区间，输出的第二座桥连接第二个区间和第三个区间，依此类推。
'''
def solve():
    n, m = ints()
    # 区间是有序的
    lr = [ints() for _ in range(n)]
    lrc = []
    for i in range(1, len(lr)):
        lrc.append((lr[i][0]-lr[i-1][1], lr[i][1]-lr[i-1][0]))
    lrc = [(x, i) for i, x in enumerate(lrc)]
    lrc.sort(key=lambda x:(x[0][0], x[0][1]))

    a = list(map(int, input().split()))
    idx = [(x, i) for i, x in enumerate(a)]
    idx.sort(key=lambda x:x[0])
    ret = [-1] * (n-1)
    if n-1 > m:
        return print("No")
    # 从短到长的桥进行遍历
    hp, j = [], 0
    for t,i in idx:
        # 选出在t>=l 的里面，r最小的数字但是>=t
        while j < len(lrc) and lrc[j][0][0] <= t:
            heappush(hp, (lrc[j][0][1], lrc[j][1]))
            j += 1
        if hp:
            if hp[0][0] < t: break
            ret[heappop(hp)[1]] = i+1

    if hp or j != len(lrc):
        return print("No")

    print("Yes")
    print(" ".join(map(str, ret)))
solve()
    
