class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        l,r = 0, 0
        add = 0
        minLen = 100000

        while r < len(nums):
            add += nums[r]
            
            if add >= target:
                
                while add >= target and l < len(nums):
                    minLen = min(r-l + 1, minLen)
                    add -= nums[l]
                    l+=1
            # update the r
            r+=1
        
        if minLen == 100000:
            return 0
        else:
            return minLen
        
        