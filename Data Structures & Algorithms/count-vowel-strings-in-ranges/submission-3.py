class Solution:
    def vowelStrings(self, words: List[str], queries: List[List[int]]) -> List[int]:
        # queries provide the range.
        # based on the range li and ri you need to check words and check if any word start or ends with a vowel

        output = []
        for query in queries:
            l, r = query[0], query[1]
            count = 0
            for i in range(l,r+1):
                if words[i][0] in "aeiou" and words[i][-1] in "aeiou":
                    count+=1
            output.append(count)
            
        return output


        