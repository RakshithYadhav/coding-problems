/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    //1. Understand    → Restate in one sentence. Make your own example.
// Check if a tree is height-balanced.
// height balanced means height of left - height of right <= 1

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


balanced, _ := helper(root)
return balanced

//9. Edge cases    → Empty, single, all same, duplicates, none/all satisfy.
}


func helper(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	leftBalanced, leftHeight := helper(root.Left)
	rightBalanced, rightHeight := helper(root.Right)
	diff := leftHeight - rightHeight
if diff < 0 {
    diff = -diff
}
	balanced := leftBalanced && rightBalanced && diff <= 1

	return balanced, 1 + max(leftHeight, rightHeight)
}
