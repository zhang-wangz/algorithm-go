# 输入 n(≤500) m(≤500) k 和一个 n 行 m 列的网格图，'#' 表示墙，'.' 表示平地。
# 保证所有 '.' 可以互相到达（四方向连通）。保证 k 小于 '.' 的个数。
# 你需要把恰好 k 个 '.' 修改成 'X'，使得剩余的所有 '.' 仍然是可以互相到达的。
# 输出修改后的网格图。
import os, sys
from io import BytesIO, IOBase

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

def solve():
    n,m,k = map(int,input().split())
    t = [input().replace('.','X') for _ in range(n)]
    k = n * m - k - sum( i.count('#') for i in t) # 所有的点数-墙数，就是平地数，-k个平地，就是需要剩余的连通
    t = [ list(i) for i in t]
    i, q = 0,[]
    # 找到一个点
    while k:
        if 'X' in t[i]:
            j = t[i].index('X')
            t[i][j],q = '.',[(i,j)]
            k -= 1
            break
        i += 1
    while k:
        x,y = q.pop()
        for i,j in ((x,y-1),(x,y+1),(x-1,y),(x+1,y)):
            if 0<=i<n and 0<=j<m and t[i][j]=='X':
                t[i][j] = '.'
                q.append((i,j))
                k -= 1
                if k == 0:
                    break
    for i in range(n):
        t[i]=''.join(t[i])
    print('\n'.join(t))

def solve2():
    n,m,k = map(int,input().split())
    t = [input().replace('.','X') for _ in range(n)]
    k = n * m - k - sum(i.count('#') for i in t) # 所有的点数-墙数，就是平地数，-k个平地，就是需要剩余的连通
    t = [ list(i) for i in t]
    i,tmp = 0,()
    # 找到一个点
    while i < n:
        if 'X' in t[i]:
            j = t[i].index('X')
            tmp = (i,j)
            break
        i += 1
    
    q = set()
    @bootstrap
    def dfs(x, y:int):
        if len(q) == k:
            yield None
        q.add((x, y))
        t[x][y] = "."
        for i,j in ((x,y-1),(x,y+1),(x-1,y),(x+1,y)):
            if 0<=i<n and 0<=j<m and (i,j) not in q and t[i][j] == 'X':
                yield dfs(i,j)
        yield None
    dfs(tmp[0], tmp[1])
    for i in range(n):
        t[i]=''.join(t[i])
    print('\n'.join(t))
solve2()
