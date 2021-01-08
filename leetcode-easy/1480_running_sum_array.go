package easy

func runningSum(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	current := nums[0]

	for i := 1; i < len(nums); i++ {
		nums[i] += current
		current = nums[i]
	}
	return nums
}
