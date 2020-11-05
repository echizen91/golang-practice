package easy

import (
	"fmt"
	"strconv"
)

/*
Given a positive integer num consisting only of digits 6 and 9.

Return the maximum number you can get by changing at most one digit (6 becomes 9, and 9 becomes 6).
*/

// convert to string
func maximum69Number(num int) int {
	numStr := fmt.Sprintf("%v", num)
	var resultStr string
	for i, v := range numStr {
		if string(v) == "6" {
			resultStr += "9" + numStr[i+1:]
			break
		}
		resultStr += string(v)

	}
	resultNum, _ := strconv.Atoi(resultStr)

	return resultNum
}
