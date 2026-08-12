func maxProfit(prices []int) int {
//1. Understand    → Restate in one sentence. Make your own example.

	// You have a list of prices. You can buy only on single day.
	// Input List of prices.
	// Output - You need to return the maximum profit you can achieve.

//2. Constraints   → What does n tell you about target complexity?
//3. Visualize     → Draw it. Boxes, pointers, trees, table — whatever fits.
//4. Brute force   → Plain English. Dumbest working solution.

// Algo
	// - Check Each index as the buy day.
	// Once you choose an index
	// The sell day should always lie between i+1, i + n
	// Note the sell day cannot be behind buy day.
	// run another loop in this range calculate profit for each index. 
	// check if the local profit is greater than current profit if its greater assign.

	n := len(prices)
	maxProfit := 0

	for i := 0; i < n; i++ {
		buy := prices[i]
		for j := i + 1; j < n; j++ {
			sell := prices[j]
			profit := sell - buy
			if profit > maxProfit {
				maxProfit = profit
			}
		} 
	}

	return maxProfit
//5. Find the waste → Where is repeated/redundant work happening?
//6. Optimize      → Name the pattern. Describe the approach in plain English.
//7. Invariant     → "At every step, X is always true." One sentence.
//8. Code          → Only now. Trace each block as you write it.
//9. Edge cases    → Empty, single, all same, duplicates, none/all satisfy.
}
