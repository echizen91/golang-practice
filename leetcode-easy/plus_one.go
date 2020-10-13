package easy

/*
Given a non-empty array of digits representing a non-negative integer, increment one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contains a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Solution:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Plus One.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for Plus One.
*/

// PlusOne : add one to numbers stored as array, will carry over if overflow
func PlusOne(digits []int) []int {
	var carryOver int
	for i := len(digits) - 1; i >= 0; i-- {
		if i == len(digits)-1 && digits[i] < 9 {
			digits[i] = digits[i] + 1
		} else if digits[i] < 9 {
			digits[i] = digits[i] + carryOver
			carryOver = 0
		} else {
			digits[i] = 0
			carryOver = 1
		}

		// Don't require to move on
		if carryOver < 1 {
			break
		}
	}

	if carryOver > 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}
