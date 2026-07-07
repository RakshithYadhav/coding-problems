class Solution:
    def replaceElements(self, arr: List[int]) -> List[int]:
        # understand
        # replace every element with the largest element to its right.
        # replace last element with -1

        # constraints.
        # len -> 10,000

        #approach -> brute force is suffcient

        output = []
        for i in range(len(arr)):
            m = -1

            for j in range(i+1, len(arr)):
                m = max(m, arr[j])
                
            output.append(m)
            print(output)
        return output       





        