class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        # We a string. we have to find the length of the longest substring  which as no duplicates.
        # Input = string s
        # output = max len of a substring in s which as no duplicates.

        # example - "" = 0
        # "ssss" = 1
        # "assavrggasfdh"= 6 "gasfdh"

        # constraints.
        # s length is 1000 max and min ""
        
        # Brute force approach 
        # Find all possible substrings, 
        # check if any substring does not contain a duplicate and then record the length
       
        # but this is a o n2 solution # the ineffciency here is.
        # i = 1 , j = 2
        #"ssbcds" if i =0, j =1 we know that substring is useless if its starts from i
        # but the current implementation still goes i = 0, j = 1,2,3,4,5, even though we know (i,j) at (0,1) 
        # will produce a duplicate substring so (0,2) does not matter

        # we have to move i until there is no duplicate and then start the process. a slide window technice.

        if len(s) == 0:
            return 0

        l = 0
        r = l+1
        se = {s[l]}
        curMax = 1
        while r < len(s):
            while s[r] in se:
                se.remove(s[l])
                l+=1
            se.add(s[r])
            curMax = max((r-l+1), curMax)
            r+=1
        return curMax
        




                

        