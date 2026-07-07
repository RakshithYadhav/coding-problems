"""
# Definition for a Node.
class Node:
    def __init__(self, val = 0, neighbors = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []
"""

class Solution:
    def cloneGraph(self, node: Optional['Node']) -> Optional['Node']:
        if not node:
            return None

        hashMap = {}
        def dfs(node):
            if node in hashMap:
                return hashMap[node]
            
            clone = Node(node.val)
            hashMap[node] = clone

            for ne in node.neighbors:
                clone.neighbors.append(dfs(ne))
            
            return clone
        return dfs(node)
        