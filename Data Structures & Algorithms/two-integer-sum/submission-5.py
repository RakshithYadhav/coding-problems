class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        mapper = {v : i for i,v in enumerate(nums)}

        for i, n in enumerate(nums):
            comp = target - n
            if comp in mapper and mapper[comp] != i:
                return [min(mapper[comp], i), max(mapper[comp], i)]
        