import sys

class Node:
    def __init__(self,data):
        self.right=self.left=None
        self.data = data
class Solution:
    def insert(self,root,data):
        if root==None:
            return Node(data)
        else:
            if data<=root.data:
                cur=self.insert(root.left,data)
                root.left=cur
            else:
                cur=self.insert(root.right,data)
                root.right=cur
        return root
    def maxDepth(self, root):
        if root is None:
            return 0
        else:
            l = self.maxDepth(root.left)
            r = self.maxDepth(root.right)
            if l > r:
                return l + 1
            else:
                return r + 1
    
    def currentLevel(self, root, level):
        if root is None:
            return
        if level == 1:
            print(root.data, end = " ")
        elif level > 1:
            self.currentLevel(root.left, level - 1)
            self.currentLevel(root.right, level - 1)
        
    def levelOrder(self, root):
        height = self.maxDepth(root)
        for i in range(1, height + 1):
            self.currentLevel(root, i)

T=int(input())
myTree=Solution()
root=None
for i in range(T):
    data=int(input())
    root=myTree.insert(root,data)
myTree.levelOrder(root)
