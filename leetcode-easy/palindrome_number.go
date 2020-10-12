package easy

import (
	"reflect"
	"strconv"
)

/*
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Follow up: Could you solve it without converting the integer to a string?

Solution Optimized:
Runtime: 8 ms, faster than 95.12% of Go online submissions for Palindrome Number.
Memory Usage: 6 MB, less than 21.85% of Go online submissions for Palindrome Number.

*/

// PalindromeNumber : check if number is palindrome (Converting to string) - Not Optimized
func PalindromeNumber(x int) bool {
	if x < 0 {
		return false
	}
	strx := strconv.Itoa(x)
	strLen := len(strx) - 1
	result := true
	start := 0

	for _, c := range strx {
		if strLen <= start {
			break
		}
		if reflect.DeepEqual(string(c), string(strx[strLen])) != true {
			result = false
			break
		}
		strLen--
		start++
	}
	return result
}

// PalindromeNumberOptimized : check if number is palindrome (Not converting to string)
func PalindromeNumberOptimized(x int) bool {
	var result = true

	if x < 0 {
		return false
	}
	// break non-negative numbers into arr
	var xArr []int
	for x > 0 {
		remainder := x % 10
		x = x / 10
		xArr = append(xArr, remainder)
	}

	for i := 0; i < len(xArr); i++ {
		if i >= len(xArr)-1-i {
			break
		}
		if xArr[i] != xArr[len(xArr)-1-i] {
			return false
		}
	}

	return result
}
