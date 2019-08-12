# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None

class Solution:
    def mergeTrees(self, t1: TreeNode, t2: TreeNode) -> TreeNode:
        if t1 == None:
            return t2
        
        stack = []
        
        self.push(stack, (t1, t2))
                
        while(len(stack) != 0):

            stack, n = self.pop(stack)
            
            if n[0]==None or n[1]==None:
                continue
            
            n[0].val += n[1].val
            if n[0].left == None:
                n[0].left = n[1].left
            else:
                self.push(stack, (n[0].left, n[1].left))
            
            
            if n[0].right == None:
                n[0].right = n[1].right
            else:
                self.push(stack, (n[0].right, n[1].right))
        
        return t1
    
    def push(self, stack, n):
        stack.append(n)
    
    def pop(self, stack):
        return stack[:-1], stack[len(stack)-1]
        
        