# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def maxDepth(self, root: TreeNode) -> int:
        if root is None:
            return 0
        
        ltree = self.maxDepth(root.left)
        rtree = self.maxDepth(root.right)
        
        if ltree > rtree:
            return ltree+1
        else:
            return rtree+1
        