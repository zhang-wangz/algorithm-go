# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def maxAncestorDiff(self, root: Optional[TreeNode]) -> int:
        q = []
        ans = 0
        def dfs(top):
            nonlocal ans
            q.append(top.val)
            if not (top.left or top.right):
                ans = max(ans, max(q)-min(q))
            if top.left:
                dfs(top.left)
            if top.right:
                dfs(top.right)
            q.pop()
        dfs(root)
        return ans