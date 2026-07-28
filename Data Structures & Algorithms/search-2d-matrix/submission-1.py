class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        # You have sorted matrix.
        # You have a target you need to search and see if it is present.
        # you need to return true or false

        # each row acts as a binary search input
        # loop through each row pass it to a helper binary search.
        def binaryHelper(row,l,r, target):
            
            while l <= r:
                m = l + ((r-l)//2)

                if row[m] < target:
                    l = m + 1
                elif row[m] > target:
                    r = m - 1
                else:
                    return True
            return False

        ROWS, COLS = len(matrix), len(matrix[0])

        l,r = 0, ROWS-1
        while l <= r:
            m = l + ((r-l) // 2)
            row = matrix[m]

            if target < row[0]:
                r = m - 1
            elif target > row[COLS-1]:
                l = m + 1
            else:
                if binaryHelper(row, 0,COLS-1, target):
                    return True
                else:
                    break
        
        return False
    
            
        