class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        adj = defaultdict(list)
        indegree = [0] * numCourses
        
        # Step 2: Build graph
        for crs, pre in prerequisites:
            adj[pre].append(crs)
            indegree[crs] += 1
            
        # Step 3: Initialize queue with courses having 0 indegree
        q = deque([i for i in range(numCourses) if indegree[i] == 0])
        order = []
        
        # Step 4: BFS Traversal (Kahn's Algorithm)
        while q:
            curr = q.popleft()
            order.append(curr)
            
            for nei in adj[curr]:
                indegree[nei] -= 1
                if indegree[nei] == 0:
                    q.append(nei)
                    
        # Step 5: Check for cycles
        if len(order) == numCourses:
            return order
        return []