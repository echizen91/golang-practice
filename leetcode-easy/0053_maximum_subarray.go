package easy

import "math"

/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Follow up:
If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

Solution:
Runtime: 4 ms, faster than 97.41% of Go online submissions for Maximum Subarray.
Memory Usage: 3.3 MB, less than 22.17% of Go online submissions for Maximum Subarray.
*/

// MaxSubArray : find maximum contiguous numbers  -> O(n) solution
func MaxSubArray(nums []int) int {
	var current, global int

	for k, val := range nums {
		if k == 0 {
			current = val
			global = val
			continue
		}
		current = int(math.Max(float64(current+val), float64(val)))
		if global < current {
			global = current
		}
	}
	return global
}
