class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        ROWS,COLS = len(grid), len(grid[0])
        count = 0

        def dfs(r,c):
            if r < 0 or c < 0 or r >= ROWS or c >= COLS or grid[r][c] != '1':
                return
            
            grid[r][c] = 'O' # processed.

            for dr,dc in [(1,0),(-1,0),(0,1),(0,-1)]:
                nr,nc = dr+r,dc+c
                dfs(nr,nc)
        
        for i in range(ROWS):
            for j in range(COLS):
                if grid[i][j] == '1':
                    # traverse to find the connected region
                    dfs(i,j)
                    count+=1
        return count