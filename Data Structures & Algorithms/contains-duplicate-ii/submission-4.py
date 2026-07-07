class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        mapper = {}
        result = False

        for i in range(len(nums)):
            if nums[i] not in mapper:
                mapper[nums[i]] = i
            else:
                prev = mapper.get(nums[i])
                mapper[nums[i]] = i
                if abs(i - prev) <= k:
                    result = True
                    
                
        
        return result
        