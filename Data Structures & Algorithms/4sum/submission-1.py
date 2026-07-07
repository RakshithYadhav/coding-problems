class Solution:
    def fourSum(self, nums: List[int], target: int) -> List[List[int]]:
        output = []
        nums.sort()
        for i in range(len(nums)):
            if i > 0 and nums[i] == nums[i-1]:
                continue
            first = nums[i]
            for j in range(i + 1, len(nums)):
                if j > i + 1 and nums[j] == nums[j-1]:
                    continue
                second = nums[j]

                l,r = j + 1, len(nums) - 1
                while l < r:
                    if first + second + nums[l] + nums[r] > target:
                        r-=1
                    elif first + second + nums[l] + nums[r] < target:
                        l+=1
                    else:
                        output.append([first, second, nums[l], nums[r]])
                        l+=1
                        r-=1
                    
                        while l < r and nums[l] == nums[l-1]:
                            l+=1
                    
                        while r > l and nums[r] == nums[r+1]:
                            r-=1
        
        return output
        