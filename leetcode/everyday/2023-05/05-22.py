# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def sufficientSubset(self, root: Optional[TreeNode], limit: int) -> Optional[TreeNode]:
        limit -= root.val
        if root.left == root.right:
            return None if limit > 0 else root
        if root.right: root.right = self.sufficientSubset(root.right, limit)
        if root.left: root.left = self.sufficientSubset(root.left, limit)
        return root if root.left or root.right else None