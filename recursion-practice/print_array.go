package recursion

import (
	"strconv"
)

/*
printing an array using recursion
*/

func printArray(nums []int) string {
	if len(nums) < 2 {
		return strconv.Itoa(nums[0])
	}

	return strconv.Itoa(nums[0]) + " " + printArray(nums[1:])
}
