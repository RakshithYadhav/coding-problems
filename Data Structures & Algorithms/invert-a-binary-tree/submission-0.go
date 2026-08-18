/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
//1. Understand    → Restate in one sentence. Make your own example.
// Inverting means reversing the values Left becomes right and right become left.
//2. Constraints   → What does n tell you about target complexity?
//3. Visualize     → Draw it. Boxes, pointers, trees, table — whatever fits.
//4. Brute force   → Plain English. Dumbest working solution.
// At each node level.
// Assign left to right and right to left.
// Then pass left and right values downstream
//5. Find the waste → Where is repeated/redundant work happening?
//6. Optimize      → Name the pattern. Describe the approach in plain English.
//7. Invariant     → "At every step, X is always true." One sentence.
// at every step left value becomes right and vice versa
//8. Code          → Only now. Trace each block as you write it.

if root == nil {
	return root
}

root.Left, root.Right = root.Right, root.Left

invertTree(root.Left)
invertTree(root.Right)

return root
//9. Edge cases    → Empty, single, all same, duplicates, none/all satisfy.
}
