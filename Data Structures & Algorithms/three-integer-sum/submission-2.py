class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        # Understand - we have a sequence of nums.
        # we have to return all possible 3 numbers that when added will sum upto 0.
        # Note - we have to return all possible which means there might be more than 1 possible solution.
        # we cant return as soon as we find a solution.
        # Note 2 - We need to store the actual numbers and not thee indexs.

        # Brute force will need 3 loops nested. o(n3) which means 1000 * 1000 * 1000 operations
        # Brute force checks each and every combination but this is slow. 
        # we know if the array was sorted left number will be lesser than right.
        # but our array is not sorted. so we do a sort operation.
        # then we go get one loop and inside that we follow the l,r pointer approach since we know the array is sorted.
        # the loop is basically for the 3rd number.

        # so there will be 2 loops instead of 3 considering the entire time complexity.
        # we can estimate it to be o(n2)

        nums.sort()
        out = []
        N = len(nums)
        for i in range(N-1):
            first = nums[i]
            if i > 0 and nums[i] == nums[i-1]:
                continue
            l = i + 1
            r = N - 1
            while l < r:
                second = nums[l]
                third = nums[r]
                total = first + second + third
                if total > 0:
                    r-=1
                elif total < 0:
                    l+=1
                else:
                    out.append([first,second,third])
                    l+=1
                    r-=1
                    while l < r and nums[l] == nums[l-1]:
                        l+=1
        return out
        