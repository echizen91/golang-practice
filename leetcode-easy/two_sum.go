package easy

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Solution for optimized:
Runtime: 4 ms, faster than 93.92% of Go online submissions for Two Sum.
Memory Usage: 3.8 MB, less than 61.67% of Go online submissions for Two Sum.
*/

// TwoSums : Sum two numbers to target (Not Optimized)
func TwoSums(nums []int, target int) []int {
	var result []int

	for i := 0; i < len(nums); i++ {
		total := nums[i]
		for j := i + 1; j < len(nums); j++ {
			if total+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
				goto end
			}
		}
	}

end:
	return result
}

// TwoSumsOptimized : Sum two numbers to target (Optimized)
func TwoSumsOptimized(nums []int, target int) []int {
	// var result []int
	numMap := make(map[int]int)

	for k, v := range nums {
		if _, ok := numMap[target-v]; ok {
			return []int{numMap[target-v], k}
		}
		numMap[v] = k

	}
	return []int{}
}
