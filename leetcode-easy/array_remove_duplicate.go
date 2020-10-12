package easy

/*
Given a sorted array nums, remove the duplicates in-place such that each element appears only once and returns the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Solution:
Runtime: 16 ms, faster than 25.57% of Go online submissions for Remove Duplicates from Sorted Array.
Memory Usage: 4.6 MB, less than 100.00% of Go online submissions for Remove Duplicates from Sorted Array.
*/

// ArrayRemoveDuplicate : Removing duplicate with in-place removal O(1) memory
func ArrayRemoveDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		currentNumLength := len(nums) - 1
		if i+1 > currentNumLength {
			break
		}
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}
