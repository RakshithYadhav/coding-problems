class Solution:
    def isPalindrome(self, s: str) -> bool:
        # A Palindrome is a string that can be read the same from back or front.

        # Input = s
        # output = boolean true if palindrome and false if not.

        # Constraints - s character is ascii
        # string len is max 1000 and min 1.

        # Brute force.
        strip = s.replace(" ","")
        print(strip.lower())
        c = ''
        for ch in strip.lower():
            if ch.isalnum():
                c+=ch
        
        print(c)
        if len(c) % 2 == 0:
            mid = len(c) // 2
        else:
            mid = (len(c) // 2 ) + 1

        print(c[:mid])
        print(c[mid:][::-1])
        return c[:len(c) // 2] == c[mid:][::-1]
        