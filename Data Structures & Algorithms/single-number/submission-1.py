class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        mask = 0
        for num in nums:
            mask = num ^ mask
        return mask
        