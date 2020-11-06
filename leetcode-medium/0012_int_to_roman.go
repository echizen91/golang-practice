package medium

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
For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral.
*/

func intToRoman(num int) string {
	var result string
	for num > 0 {
		if num >= 1000 {
			result += "M"
			num -= 1000
		} else if num >= 900 {
			result += "CM"
			num -= 900
		} else if num >= 500 {
			result += "D"
			num -= 500
		} else if num >= 400 {
			result += "CD"
			num -= 400
		} else if num >= 100 {
			result += "C"
			num -= 100
		} else if num >= 90 {
			result += "XC"
			num -= 90
		} else if num >= 50 {
			result += "L"
			num -= 50
		} else if num >= 40 {
			result += "XL"
			num -= 40
		} else if num >= 10 {
			result += "X"
			num -= 10
		} else if num >= 9 {
			result += "IX"
			num -= 9
		} else if num >= 5 {
			result += "V"
			num -= 5
		} else if num >= 4 {
			result += "IV"
			num -= 4
		} else {
			result += "I"
			num--
		}
	}
	return result
}
