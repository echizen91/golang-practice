package easy

/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two one's added together.
12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right.
However, the numeral for four is not IIII. Instead, the number four is written as IV.
Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX.
There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.

Given a roman numeral, convert it to an integer.

Solution:
Runtime: 4 ms, faster than 91.58% of Go online submissions for Roman to Integer.
Memory Usage: 3.1 MB, less than 8.70% of Go online submissions for Roman to Integer.
*/

// RomanToInt : convert roman numeral to integers
func RomanToInt(s string) int {
	var val int
	for i := 0; i < len(s); i++ {
		firstChar := string(s[i])
		var nextChar string
		if i+1 < len(s) {
			nextChar = string(s[i+1])
		}
		/*
			Cover Edge Cases
		*/
		if firstChar == "I" && nextChar == "V" {
			val += 4
			i++
			continue
		}
		if firstChar == "I" && nextChar == "X" {
			val += 9
			i++
			continue
		}
		if firstChar == "X" && nextChar == "L" {
			val += 40
			i++
			continue
		}
		if firstChar == "X" && nextChar == "C" {
			val += 90
			i++
			continue
		}
		if firstChar == "C" && nextChar == "D" {
			val += 400
			i++
			continue
		}
		if firstChar == "C" && nextChar == "M" {
			val += 900
			i++
			continue
		}
		if firstChar == "M" {
			val += 1000
			continue
		}
		if firstChar == "D" {
			val += 500
			continue
		}
		if firstChar == "C" {
			val += 100
			continue
		}
		if firstChar == "L" {
			val += 50
			continue
		}
		if firstChar == "X" {
			val += 10
			continue
		}
		if firstChar == "V" {
			val += 5
			continue
		}
		if firstChar == "I" {
			val++
			continue
		}

	}
	return val
}
