func findMin(nums []int) int {
//1. Understand    → Restate in one sentence. Make your own example.
	// You have a sorted array which is rotated. 
	// You need to find the smallest element in the array
	// All the elements in the array are unique.
//2. Constraints   → What does n tell you about target complexity?
	// Can you find the o(log n) solution
//3. Visualize     → Draw it. Boxes, pointers, trees, table — whatever fits.
//1,2,3,4,5,6    R = 1
// 6,1,2,3,4,5 5 / 2

// 9, 1, 2, 3, 4, 5, 6, 7 , 8

// mid = 2
// r = 2

//4. Brute force   → Plain English. Dumbest working solution.

// Find the partition will give use the largest and small.
// In a sorted array. all elements will follow this rule.

// result := nums[0]
// for index := range nums {
// 	if index == 0 {
// 		continue
// 	}

// 	if nums[index] < nums[index-1] {
// 		result = nums[index]
// 	}
// }

// return result

//5. Find the waste → Where is repeated/redundant work happening?
	// You make the comparison at each index.
//6. Optimize      → Name the pattern. Describe the approach in plain English.

// 1, 2, 3
// 3, 1, 2

// 1, 2, 3, 4
// 4, 1, 2, 3

// In a sorted array.
// mid will always be greater than current l
// mid will always be lesser than current r

// Reason the array is sorted.

// In a rotated array.
// mid can be lesser than current l.

//7. Invariant     → "At every step, X is always true." One sentence.
// At every step, the minimum always lies within [l, r]

//8. Code          → Only now. Trace each block as you write it.
	l, r := 0, len(nums) - 1

	for l < r {
		mid := l + (r-l) / 2
		// array as been rotated.
		if nums[mid] > nums[r] {
			l = mid + 1	
		} else {
			 r = mid
		}
	}

	return nums[l]
//9. Edge cases    → Empty, single, all same, duplicates, none/all satisfy.
}
