class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        # Understand
        # Find and group similar anagrams.

        # constraints are small so a brute force is good.

        # bruteforce approach 
        # use a map the key is the sorted(str) and values is a list of string 
        mapper = defaultdict(list)
        for string in strs:
            key = "".join(sorted(string))
            print(key)
            mapper[key].append(string)
    
        output = []
        for k in mapper:
            output.append(mapper[k])
        
        return output
            
        