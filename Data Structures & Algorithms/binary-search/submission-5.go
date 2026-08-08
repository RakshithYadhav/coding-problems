func search(nums []int, target int) int {
//1. Understand    → Restate in one sentence. Make your own example.
	// Array with No duplicate Sorted So basically we have to find the given target.
	//Output is basically returning the index of the element that we are searching.
//2. Constraints   → What does n tell you about target complexity?
	// Constance is basically. So we have 10,000. But the algorithm itself should be O of N log n
//3. Visualize     → Draw it. Boxes, pointers, trees, table — whatever fits.
	// Since Ar Sorted You can find the mid And cheque if the maid is equal to the target If the mid is not equal to the target, and if mid is less than target, Then we have to cheque the left side, which obviously is lesser than the target. So then you would have to update the right pointer If the mid is greater than the target, then it means the target elements lies to the right side of the mat, because it is array sorted. So in that case, we have to cheque the right part So we have to update the left to be mid one. So this is the general algorithm.
//4. Brute force   → Plain English. Dumbest working solution.
	l := 0
	r := len(nums) - 1 

	for l <= r {
		mid := l + (r - l) / 2
		fmt.Println(mid)
		if nums[mid] < target {
			l = mid + 1
		}else if nums[mid] > target {
			r = mid - 1
		}else {
			return mid
		}
	}
	return -1
//5. Find the waste → Where is repeated/redundant work happening?
//6. Optimize      → Name the pattern. Describe the approach in plain English.
//7. Invariant     → "At every step, X is always true." One sentence.
//8. Code          → Only now. Trace each block as you write it.
//9. Edge cases    → Empty, single, all same, duplicates, none/all satisfy.
}
