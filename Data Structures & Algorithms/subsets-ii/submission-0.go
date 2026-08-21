func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make(map[string][]int)

	var backtrack func(int, []int)
	backtrack = func(i int, subset []int) {
		if i == len(nums) {
			key := fmt.Sprint(subset)
			res[key] = append([]int{}, subset...)
			return
		}

		subset = append(subset, nums[i])
		backtrack(i+1, subset)
		subset = subset[:len(subset)-1]
		backtrack(i+1, subset)
	}

	backtrack(0, []int{})

	var result [][]int
	for _, v := range res {
		result = append(result, v)
	}
	return result
}
