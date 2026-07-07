class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        count1 = Counter(nums)
        for n in nums:
            if count1[n] > 1:
                return True
        return False
        