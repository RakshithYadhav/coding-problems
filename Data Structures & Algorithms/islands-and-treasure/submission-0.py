class Solution:
    def islandsAndTreasure(self, grid: List[List[int]]) -> None:
        ROWS, COLS = len(grid), len(grid[0])
        directions = [(0,1),(0,-1),(1,0),(-1,0)]
        visited = set()
        inf = 2147483647


        queue = collections.deque()
        for i in range(ROWS):
            for j in range(COLS):
                if grid[i][j] == 0:
                    queue.append((i,j))
                    visited.add((i,j))
        
        while queue:
            r,c = queue.popleft()
            for dr,dc in directions:
                nr , nc = dr + r , dc + c
                if (0 <= nr < ROWS and 0 <= nc < COLS and grid[nr][nc] == inf):
                    grid[nr][nc] = 1 + grid[r][c]
                    queue.append((nr,nc))





        