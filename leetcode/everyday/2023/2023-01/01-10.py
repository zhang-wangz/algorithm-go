# 求解欧拉回路 / 欧拉通路的题目 753
# https://leetcode.cn/problems/cracking-the-safe/solution/po-jie-bao-xian-xiang-by-leetcode-solution/
class Solution:
    def crackSafe(self, n: int, k: int) -> str:
        mod = 10 ** (n-1)
        ans = list()
        vis = set()
        def dfs(node):
            for i in range(k):
                v = node * 10 + i
                if v in vis: continue
                vis.add(v) 
                dfs(v % mod) # 节点编号, 去除最高位
                ans.append(str(i)) # 保留最后一个数
        dfs(0)
        return "".join(ans) + "0" * (n-1) # 最后加上0的前缀0数量