func lengthOfLongestSubstring(s string) int {
//1. Understand    → Restate in one sentence. Make your own example.
	// Returns the length of the longest substring Inside the mainstream. Without Duplicate.
//2. Constraints   → What does n tell you about target complexity?
//3. Visualize     → Draw it. Boxes, pointers, trees, table — whatever fits.
//4. Brute force   → Plain English. Dumbest working solution.
// Brute Force is to use 2 loops check all substrings and check if they are duplicate or not and record the length.
//5. Find the waste → Where is repeated/redundant work happening?
//6. Optimize      → Name the pattern. Describe the approach in plain English.
//7. Invariant     → "At every step, X is always true." One sentence.
//Invariant - No Duplicate should exist between L and R
//8. Code          → Only now. Trace each block as you write it.
// use a set
unique := make(map[byte]struct{})
l,r := 0,0
maxLen := 0

for r < len(s) {
	for {
		_, exists := unique[s[r]]
		if !exists {
			break
		}
    	delete(unique, s[l])
    	l++
	}

	cur := (r-l+1)

	if cur > maxLen {
		maxLen = cur
	}

	unique[s[r]] = struct{}{}
	r++
}

return maxLen

//9. Edge cases    → Empty, single, all same, duplicates, none/all satisfy.
}
