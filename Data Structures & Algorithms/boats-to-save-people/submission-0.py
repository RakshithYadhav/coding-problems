class Solution:
    def numRescueBoats(self, people: List[int], limit: int) -> int:
        people.sort()
        output = 0
        l,r = 0,  len(people)-1

        while l <= r:
            totalWeight = people[l] + people[r]
            if totalWeight <= limit:
                output+=1
                l+=1
                r-=1
            else:
                output+=1
                r-=1
        return output