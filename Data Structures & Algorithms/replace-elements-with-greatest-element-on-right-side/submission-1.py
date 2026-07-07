class Solution:
    def replaceElements(self, arr: List[int]) -> List[int]:
        # understand
        # replace every element with the largest element to its right.
        # replace last element with -1

        # constraints.
        # len -> 10,000

        #approach -> brute force is suffcient 
        n = len(arr)
        ans = [0] * n
        rightMax = -1
        for i in range(n-1,-1,-1):
            ans[i] = rightMax
            rightMax = max(arr[i], rightMax)
        return ans  





        