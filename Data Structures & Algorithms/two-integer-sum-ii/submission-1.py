class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        # Understand - We have a array of numbers which is sorted in non decreasing order.
        # basically everything keeps increasing or remains the same
        # [1,2,3,4]

        # we also have a target. 
        # Task - return the indices of those numbers in the array (1 index based) which when summed together 
        # will add upto the target.

        ## [3,4,7,8] = Target = 11
        ## 3, 8 will add upto the target 3 index is 1 and 8 is 4
        ## [1,4]

        # Constraints array lenght is quite small a simple brute force solution will do the trick.

        for i in range(len(numbers)):
            tempA = numbers[i]
            for j in range(i+1, len(numbers)):
                tempB = numbers[j]
                if tempA + tempB == target:
                    return [i+1, j+1]
                    

        
                
        