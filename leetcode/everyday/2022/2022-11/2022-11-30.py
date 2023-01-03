from collections import Counter


class FreqStack:

    def __init__(self):
        self.cnt = Counter()
        self.st = []

    def push(self, val: int) -> None:
        if self.cnt[val] == len(self.st):
            self.st.append([val])
        else:
            self.st[self.cnt[val]].append(val)
        self.cnt[val] += 1
    

    def pop(self) -> int:
        val = self.st[-1].pop()
        if len(self.st[-1]) == 0:
            self.st.pop()
        self.cnt[val] -= 1
        return val


# Your FreqStack object will be instantiated and called as such:
# obj = FreqStack()
# obj.push(val)
# param_2 = obj.pop()