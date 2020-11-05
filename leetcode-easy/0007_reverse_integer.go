package easy

import (
	"math"
)

/*
Given a 32-bit signed integer, reverse digits of an integer.

Note:
Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

Solution:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Reverse Integer.
Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Reverse Integer.


*/

// ReverseInteger : reverse the order of integer (Optimized)
func ReverseInteger(x int) int {
	var negative bool
	if x < 0 {
		negative = true
	}

	x = int(math.Abs(float64(x)))
	result := 0
	for x > 0 {
		remainder := x % 10
		result = result*10 + remainder
		x /= 10
	}

	if result > math.MaxInt32 {
		result = 0
	}

	if negative {
		result *= -1
	}

	return result
}
