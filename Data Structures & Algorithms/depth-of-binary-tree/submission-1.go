/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
//1. Understand    → Restate in one sentence. Make your own example.
// return the maximum depth the number of nodes along the longest path.

//2. Constraints   → What does n tell you about target complexity?
//3. Visualize     → Draw it. Boxes, pointers, trees, table — whatever fits.
//4. Brute force   → Plain English. Dumbest working solution.
//5. Find the waste → Where is repeated/redundant work happening?
//6. Optimize      → Name the pattern. Describe the approach in plain English.
//7. Invariant     → "At every step, X is always true." One sentence.
//8. Code          → Only now. Trace each block as you write it.
if root == nil {
	return 0
}

return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
//9. Edge cases    → Empty, single, all same, duplicates, none/all satisfy.
}
