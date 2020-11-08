package hard

/*
Given an array of non-negative integers nums, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.

You can assume that you can always reach the last index.
*/

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func jump(nums []int) int {
	var currentIndex, lastIndex, steps int

	for i := 0; i < len(nums); i++ {
		if i > lastIndex {
			lastIndex = currentIndex
			steps++
		}
		currentIndex = max(currentIndex, nums[i]+i)
	}

	return steps
}
