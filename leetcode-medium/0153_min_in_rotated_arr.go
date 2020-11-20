package medium

func findMin(nums []int) int {
	left := nums[0]
	right := nums[len(nums)-1]
	// sorted array
	if left < right {
		return left
	}
	for i := 1; i <= len(nums)/2; i++ {
		if left > nums[i] {
			return nums[i]
		}
		if right < nums[len(nums)-i-1] {
			return right
		}
		left = nums[i]
		right = nums[len(nums)-i-1]
	}
	return left
}
