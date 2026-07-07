class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        for i, n in enumerate(numbers):
            comp = target - n
            for j , m in enumerate(numbers):
                if m == comp:
                    return [i+1, j+1]
                
        