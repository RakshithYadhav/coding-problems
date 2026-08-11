func search(nums []int, target int) int {
//1. Understand    → Restate in one sentence. Make your own example.
	// You have a sorted array which is rotated.
	// You have a target element which you have to find.

//2. Constraints   → What does n tell you about target complexity?
//3. Visualize     → Draw it. Boxes, pointers, trees, table — whatever fits.
	// 3,4,5,6,7,8,9, 10, 1,2

	//6,1,2,3,4,5 

	// target = 1,2,3,4,5,6



	// 4, 5, 6, 1,2,3 Target equal  = 1,2,3
	// nums[mid] > nums[r]
		// l,mid is sorted.
	// nums[mid] < nums[r]
		// mid to r sorted.

		

//4. Brute force   → Plain English. Dumbest working solution.
	// a rotated sorted array always produces one sorted part and one unsorted part.
	// First check which is sorted and which is unsorted.
	// next determine if the element exists inside the sorted or unsorted.
	// if the element exists in the sorted do a binary search algorithm.
	// if the element exists insdie the unsorted sicne you already eliminated half do a trivial search

	l,r := 0, len(nums) - 1
	mid := (l + r) / 2
	isInSorted := false
	var sortedLeft, sortedRight, unsortedLeft, unsortedRight int

	if nums[mid] > nums[r] {
		unsortedLeft, unsortedRight = mid, r
		sortedLeft, sortedRight = l, mid
	} else {
		unsortedLeft, unsortedRight = l, mid
		sortedLeft, sortedRight = mid, r
	}

	// the element is inside the sorted range perform binary search.
	if nums[sortedLeft] <= target && target <= nums[sortedRight] {
		isInSorted = true
	}

	if isInSorted {
		// perform a binary search.
		left, right := sortedLeft, sortedRight
		for left <= right {
			m := (left + right) / 2

			if target < nums[m] {
				right = m - 1
			} else if target > nums[m] {
				left = m + 1
			} else {
				return m
			}
		}
		return -1

	}else {
		// go through the unsorted range and find the index.
		for index := unsortedLeft ; index <= unsortedRight; index ++ {
			if nums[index] == target {
				return index
			}
		}
		return -1
	}
	
//5. Find the waste → Where is repeated/redundant work happening?
//6. Optimize      → Name the pattern. Describe the approach in plain English.
//7. Invariant     → "At every step, X is always true." One sentence.
//8. Code          → Only now. Trace each block as you write it.
//9. Edge cases    → Empty, single, all same, duplicates, none/all satisfy.
}
