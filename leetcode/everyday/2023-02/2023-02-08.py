class trie():
    def __init__(self):
        self.child = dict()
        self.ref = -1
# 字典树
class Solution:
    def removeSubfolders(self, folder: List[str]) -> List[str]:
        root = trie()
        for i, path in enumerate(folder):
            t = root
            path = path.split('/')
            for name in path:
                if name not in t.child:
                    t.child[name] = trie()
                t = t.child[name]
            t.ref = i
        ans = []
        def dfs(t: trie):
            if t.ref != -1:
                ans.append(folder[t.ref])
                return
            for c in t.child.values():
                dfs(c)
        dfs(root)
        return ans
# sort+map
class Solution:
    def removeSubfolders(self, folder: List[str]) -> List[str]:
        folder.sort()
        ans = [folder[0]]
        for i in range(1, len(folder)):
            if not ((pre := len(ans[-1])) < len(folder[i]) and ans[-1] == folder[i][:pre] and folder[i][pre] == '/'):
                ans.append(folder[i])
        return ans