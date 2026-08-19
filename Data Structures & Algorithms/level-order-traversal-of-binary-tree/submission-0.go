/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
//1. Understand    → Restate in one sentence. Make your own example.
	// Level Order Traversal basically means - Queue. Level by level which mean BFS search
	// BFS search means using queue.
//2. Constraints   → What does n tell you about target complexity?

//3. Visualize     → Draw it. Boxes, pointers, trees, table — whatever fits.
//4. Brute force   → Plain English. Dumbest working solution.
//5. Find the waste → Where is repeated/redundant work happening?
//6. Optimize      → Name the pattern. Describe the approach in plain English.
//7. Invariant     → "At every step, X is always true." One sentence.
//8. Code          → Only now. Trace each block as you write it.
	var out [][]int
	var queue = []*TreeNode{root}
	if root == nil {
		return out
	}

	for len(queue) > 0 {
		length := len(queue)
		level := make([]int, 0, length)

		for i := 0; i < length; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		out = append(out, level)
	}

	return out


	
//9. Edge cases    → Empty, single, all same, duplicates, none/all satisfy.
}
