package easy

/*
Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Solution:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Sqrt(x).
Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Sqrt(x).
*/

// SquareRoot : implement the square root function (return whole number only) - Newton's Integer
func SquareRoot(x int) int {
	r := x
	for r*r > x {
		r = (r + (x / r)) / 2
	}
	return r
}
