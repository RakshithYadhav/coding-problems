class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        # You have a array of numbers.
        # You have a k integer.

        # Output : K most frequently appeared integers.
        # Example 1 : [3,4,4,1,2,3,7,8] , k = 3
        # Find 3 most frequently appeared integers in the above array.

        # Most frequency Find the number which appears the most. that will be the first number.
        # Then do the same process until you reach k numbers.
        heap = []
        count = defaultdict(int)
        
        for num in nums:
            count[num]+=1
           
        sortted = sorted(count.items(), key=lambda item: item[1], reverse=True)

        for num, count in sortted:
            if k <= 0:
                break
            heap.append(num)
            k-=1
        
        return heap