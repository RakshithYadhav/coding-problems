func subsets(nums []int) [][]int {
//1. Understand    → Restate in one sentence. Make your own example.
	// 
//2. Constraints   → What does n tell you about target complexity?
//3. Visualize     → Draw it. Boxes, pointers, trees, table — whatever fits.
//4. Brute force   → Plain English. Dumbest working solution.
//5. Find the waste → Where is repeated/redundant work happening?
//6. Optimize      → Name the pattern. Describe the approach in plain English.
//7. Invariant     → "At every step, X is always true." One sentence.
//8. Code          → Only now. Trace each block as you write it.
	res := [][]int{}
	subset := []int{}

	var dfs func(int)
	dfs = func(i int) {
		if i >= len(nums) {
			temp := make([]int, len(subset))
			copy(temp, subset)
			res = append(res, temp)
			return
		}
		subset = append(subset, nums[i])
		dfs(i+1)
		subset = subset[:len(subset)-1]
		dfs(i+1)
	}
	dfs(0)
	return res
//9. Edge cases    → Empty, single, all same, duplicates, none/all satisfy.
}
