package medium

/*
Given an array nums of n integers where n > 1,  return an array output such that output[i] is equal to the product of all the elements of nums except nums[i].
Constraint: It's guaranteed that the product of the elements of any prefix or suffix of the array (including the whole array) fits in a 32 bit integer.

Note: Please solve it without division and in O(n).
*/

func productExceptSelf(nums []int) []int {
	var result []int

	// first pass -> left
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result = append(result, 1)
			continue
		}

		result = append(result, nums[i-1]*result[i-1])
	}

	// second pass <- right
	seed := 1
	for i := len(nums) - 2; i >= 0; i-- {
		seed *= nums[i+1]
		result[i] *= seed
	}

	return result
}
