class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        minLen = min(len(word1), len(word2))
        output = ""

        for i in range(minLen):
            output += word1[i] + word2[i]
        
        if len(word1) > minLen:
            output += word1[minLen:]
        else:
            output += word2[minLen:]
        
        return output     
        
        