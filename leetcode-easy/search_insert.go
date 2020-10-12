package easy

/*
Given a sorted array of distinct integers and a target value, return the index if the target is found.
If not, return the index where it would be if it were inserted in order.

Solution:
Runtime: 4 ms, faster than 81.08% of Go online submissions for Search Insert Position.
Memory Usage: 3.1 MB, less than 100.00% of Go online submissions for Search Insert Position.
*/

// SearchInsert :
func SearchInsert(nums []int, target int) int {
	var result int
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			result = i
			break
		}
		if i == len(nums)-1 && nums[i] < target {
			result = i + 1
		}
	}
	return result
}
