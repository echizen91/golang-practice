package easy

/*
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Solution:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Add Binary.
Memory Usage: 2.6 MB, less than 54.29% of Go online submissions for Add Binary.
*/

// AddBinary : Add two string binary
func AddBinary(a string, b string) string {
	var short, long, result string
	if len(a) >= len(b) {
		long = a
		short = b
	} else {
		short = a
		long = b
	}

	maintainShort := len(short) - 1
	var carryOver, stop int
	for i := len(long) - 1; i >= 0; i-- {
		if i == 0 {
			stop = -1
		}
		if maintainShort < 0 {
			stop = i
			break
		}

		if string(long[i]) == string(short[maintainShort]) {
			// Both 1, Carry Over
			if string(long[i]) == "1" {
				if carryOver == 1 {
					result = "1" + result
				} else {
					result = "0" + result
					carryOver++
				}

			} else {
				if carryOver == 0 {
					result = "0" + result
				} else {
					result = "1" + result
					carryOver--
				}
			}
		} else {
			if carryOver == 0 {
				result = "1" + result
			} else {
				result = "0" + result
			}
		}
		maintainShort--
	}

	for i := stop; i >= 0; i-- {
		if string(long[i]) == "1" {
			if carryOver == 1 {
				result = "0" + result
			} else {
				result = "1" + result
				carryOver = 0
			}
		} else {
			if carryOver == 1 {
				result = "1" + result
				carryOver = 0
			} else {
				result = "0" + result
			}
		}
	}
	if carryOver == 1 {
		result = "1" + result
	}

	return result
}
