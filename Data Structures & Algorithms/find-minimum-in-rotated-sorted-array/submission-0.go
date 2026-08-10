func findMin(nums []int) int {
//1. Understand    → Restate in one sentence. Make your own example.
	// You have a sorted array which is rotated. 
	// You need to find the smallest element in the array
	// All the elements in the array are unique.
//2. Constraints   → What does n tell you about target complexity?
	// Can you find the o(log n) solution
//3. Visualize     → Draw it. Boxes, pointers, trees, table — whatever fits.
//1,2,3,4,5,6    R = 1
// 6,1,2,3,4,5



//4. Brute force   → Plain English. Dumbest working solution.

// Find the partition will give use the largest and small.
// In a sorted array. all elements will follow this rule.

result := nums[0]
for index := range nums {
	if index == 0 {
		continue
	}

	if nums[index] < nums[index-1] {
		result = nums[index]
	}
}

return result

//5. Find the waste → Where is repeated/redundant work happening?
//6. Optimize      → Name the pattern. Describe the approach in plain English.
//7. Invariant     → "At every step, X is always true." One sentence.
//8. Code          → Only now. Trace each block as you write it.
//9. Edge cases    → Empty, single, all same, duplicates, none/all satisfy.
}
