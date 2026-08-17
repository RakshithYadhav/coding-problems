
//1. Understand    → Restate in one sentence. Make your own example.
	// A string which as only upper case characters.
	// a integer K times you can replace any Character in the string.
	// After performing a maximum K replacements, return the length of longest string which just as one character. [All characters are the same]

	// Return the length of Longest substring with all the same characters , You can do K replacements in the orginal string.

//2. Constraints   → What does n tell you about target complexity?
//3. Visualize     → Draw it. Boxes, pointers, trees, table — whatever fits.
//4. Brute force   → Plain English. Dumbest working solution.
//5. Find the waste → Where is repeated/redundant work happening?
//6. Optimize      → Name the pattern. Describe the approach in plain English.
//7. Invariant     → "At every step, X is always true." One sentence.
// Invariant - In range l,r the max length of single character is 
// Lenght of the range - highest frequence count <= k
//Lenght of the range - lowest frequence count <= k
// Length of the range - lowest frequence count <= k
//8. Code          → Only now. Trace each block as you write it.
func characterReplacement(s string, k int) int {
    count := make(map[byte]int)
    res, l, maxf := 0, 0, 0

    for r := 0; r < len(s); r++ {
        count[s[r]]++
        if count[s[r]] > maxf {
            maxf = count[s[r]]
        }

        for (r - l + 1) - maxf > k {
            count[s[l]]--
            l++
        }

        if r - l + 1 > res {
            res = r - l + 1
        }

   
	}
	return res
}
//9. Edge cases    → Empty, single, all same, duplicates, none/all satisfy.


