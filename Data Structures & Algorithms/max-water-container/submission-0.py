class Solution:
    def maxArea(self, heights: List[int]) -> int:
        # Problem - You have a set of heights You need to find any two bars which can produce the greatest area.

        # Example - [1,2,3,4,5,6,7]
        # area of a rectangle container  = w * l 
        # in order to find the max we need to make sure w * l is as great as possible.
        # We know l is the greatest between the first bar and last bar. so we can set a pointer l = 0 and r = len(heights) - 1
        # as for width we need to find min(heights[l], heights[r]) because that is the limit so water can stay inside the container.
        # if heights[l] < heights[r] then there may be a bar which is better than current l position so increase.
        # if heights[r] < heights[l] then there may be a bar which is better than current r posisition.

        l = 0
        r = len(heights)-1
        m = 0
        while l < r:
            area = min(heights[l],heights[r]) * (r-l)
            m = max(m, area)

            if heights[l] < heights[r]:
                l+=1
            elif heights[r] < heights[l]:
                r-=1
            else:
                l+=1
                r-=1
        
        return m
             
        