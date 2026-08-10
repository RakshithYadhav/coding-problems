func minEatingSpeed(piles []int, h int) int {
//1. Understand    → Restate in one sentence. Make your own example.
// You have a pile of bananans and have h before which you have to finish eating all the bananas.
// You can choose a time k which is the eating rate per hour.
//2. Constraints   → What does n tell you about target complexity?
// If you finish a pile within h meaning if pile[i] < k then you cannot immediately move to another pile. You have to wait.
//3. Visualize     → Draw it. Boxes, pointers, trees, table — whatever fits.
//4. Brute force   → Plain English. Dumbest working solution.
	// If you could have done this it would be good.
	// Simple but you directly wanted a optimized version.
	// Algorithm
		// Start with speed 1.
		// Calculate the total time with this speed.
		// if total time < h return speed
		// else increase speed and repeat
		// speed := 1
		// for {
		// 	totalHours := 0
		// 	for _, pile := range(piles) {
		// 		totalHours +=  							       		int(math.Ceil(float64(pile)/float64(speed)))
		// 	}

		// 	if totalHours <= h {
		// 		return speed
		// 	}
		// 	speed += 1
		// }
		// return speed
//5. Find the waste → Where is repeated/redundant work happening?
// Checks each speed one by one.
//6. Optimize      → Name the pattern. Describe the approach in plain English.

//7. Invariant     → "At every step, X is always true." One sentence.
// Total Time Needed Decreases as the eating speed increases.
//8. Code          → Only now. Trace each block as you write it.
	l,r := 1, 0
	for _, p := range piles {
		if p > r {
			r = p
		}
	}
	res := r

	for l <= r {
		k := (l+r) / 2
		totalTime := 0

		for _, p := range piles {
			totalTime += int(math.Ceil(float64(p) / float64(k)))
		}

		if totalTime <= h {
			res = k
			r = k - 1
		} else {
			l = k + 1
		}
	}
	return res
//9. Edge cases    → Empty, single, all same, duplicates, none/all satisfy.
}
