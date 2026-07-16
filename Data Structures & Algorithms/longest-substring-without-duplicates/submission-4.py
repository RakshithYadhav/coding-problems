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
        maxLen = 1
        for i in range(len(s)):
            for j in range(i+1, len(s)):
                temp = s[i:j+1]
                se = set(temp)
                if len(se) != len(temp):
                    continue
                maxLen = max(maxLen, len(temp))
        
        return maxLen if s else 0

                

        