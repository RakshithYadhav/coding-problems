type MinStack struct {
	stack []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack {
		stack : []int{},
		minStack : []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	minVal := val

	if len(this.minStack) > 0 {
		top := this.minStack[len(this.minStack)-1]
		if top < val {
			minVal = top
		}
	}

	this.minStack = append(this.minStack, minVal)
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

//1. Understand    → Restate in one sentence. Make your own example.
	// Failed to understand the specification of pop.
//You need to fill The implementation for a min stack. So basically what a min stack is. It //is similar to a stack where you push the elements and you pop the elements, and you get //the top element. And But it also can the data structure also can get you the Minimum //element in the in the stack. So it also is similar to a min E kind of data structure
//2. Constraints   → What does n tell you about target complexity?
//3. Visualize     → Draw it. Boxes, pointers, trees, table — whatever fits.
//4. Brute force   → Plain English. Dumbest working solution.
//5. Find the waste → Where is repeated/redundant work happening?
//6. Optimize      → Name the pattern. Describe the approach in plain English.
//7. Invariant     → "At every step, X is always true." One sentence.
//8. Code          → Only now. Trace each block as you write it.
//9. Edge cases    → Empty, single, all same, duplicates, none/all satisfy.
