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

        
        


        # what is the inefficiency here ? 
        # we check all combinations when we know the array is sorted, and some options can be removed.
        # for example
        #[1,2,3,4] = target = 3
        #i = 1
        # j = 2
        # in this case j was in the best place. Now lets evaluate the worst case.

        #[1,1,1,1,2,1,1,1,......,10]  target = 12 N = 1000
        #Now I = 1 and j = 1 and it goes until 1000 just to fail.

        # can we improve this. instead of going in a linear fashion.
        # Since the array is sorted we know the element to left is always less than right.
        # so take one low and one high. basically have a 2 pointer approach.

        l,r = 0, len(numbers)-1

        while l < r:
            s = numbers[l] + numbers[r]

            # 3 possible cases for s. greater, less or equal.
            # greater
            if s > target:
                r-=1
            elif s < target:
                l+=1
            else:
                return [l+1, r+1]


        
                
        