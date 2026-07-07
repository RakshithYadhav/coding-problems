class Solution:
    def validTree(self, n: int, edges: List[List[int]]) -> bool:
        if len(edges) != n - 1:
            return False
            
        parent = list(range(n))
        rank = [0] * n

        def find(x):
            if parent[x] != x:
                parent[x] = find(parent[x])
            return parent[x]
        
        def union(u,v):
            rootX, rootY = find(u), find(v)

            if rootX != rootY:
                if rank[rootX] > rank[rootY]:
                    parent[rootY] = rootX
                elif rank[rootY] > rank[rootX]:
                    parent[rootX] = rootY
                else:
                    parent[rootY] = rootX
                    rank[rootX]+=1
                return True
            return False
        
        for u,v in edges:
            if not union(u,v):
                return False
        return True
        
        