class Solution:
    def findRedundantConnection(self, edges: List[List[int]]) -> List[int]:
        parent = {}
        output = []

        def add(x):
            if x not in parent:
                parent[x] = x

        def find(x):
            add(x)
            if parent[x] != x:
                parent[x] = find(parent[x])
            return parent[x]
        
        def union(x,y):
            rootX, rootY = find(x), find(y)
            # we found a cycle 
            if parent[x] == parent[y]:
                output.append([x,y])
                return

            parent[rootY] = rootX
        
        for u,v in edges:
            union(u,v)
        
        return output.pop()

        
        