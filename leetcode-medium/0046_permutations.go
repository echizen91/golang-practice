package medium

/*
Given a collection of distinct integers, return all possible permutations.
*/

var result [][]int

func recursive(nums []int, length int) {
	// unique solution found
	if length == len(nums) {
		result = append(result, append(make([]int, 0), nums...))
		return
	}

	for j := length; j < len(nums); j++ {
		// swap
		nums[j], nums[length] = nums[length], nums[j]
		// call permute again
		recursive(nums, length+1)
		// backtrack
		nums[j], nums[length] = nums[length], nums[j]
	}
}

func permute(nums []int) [][]int {
	result = make([][]int, 0)
	recursive(nums, 0)
	return result
}
