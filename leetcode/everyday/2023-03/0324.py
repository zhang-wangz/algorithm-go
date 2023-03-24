class trie:
    def __init__(self):
        self.child = [None]*26
        self.isEnd = False

    def insert(self, word:str):
        tr = self
        for w in word[::-1]:
            c = ord(w) - ord('a')
            if tr.child[c] == None:
                tr.child[c] = trie()
            tr = tr.child[c]
        tr.isEnd = True

    def search(self, word:List[str])->bool:
        tr = self
        for w in word[::-1]:
            c = ord(w) - ord('a')
            if not tr.child[c]:
                return False
            tr = tr.child[c]
            if tr.isEnd:
                return True
        return False

class StreamChecker:

    def __init__(self, words: List[str]):
        self.tr = trie()
        self.limit = 200
        self.cs = []
        for w in words:
            self.tr.insert(w)
        

    def query(self, letter: str) -> bool:
        self.cs.append(letter)
        return self.tr.search(self.cs[-self.limit:])



# Your StreamChecker object will be instantiated and called as such:
# obj = StreamChecker(words)
# param_1 = obj.query(letter)