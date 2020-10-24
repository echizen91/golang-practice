package easy

import (
	"math/bits"
)

/*
Given a positive integer num, output its complement number. The complement strategy is to flip the bits of its binary representation.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Number Complement.
Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Number
*/
func findComplement(num int) int {
	return num ^ (1<<(bits.Len(uint(num))) - 1)
}
