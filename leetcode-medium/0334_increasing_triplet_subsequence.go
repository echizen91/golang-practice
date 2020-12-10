package medium

import "math"

func increasingTriplet(nums []int) bool {
	small, medium := math.MaxInt32, math.MaxInt32
	for _, val := range nums {
		if val <= small {
			small = val
		} else if val <= medium {
			medium = val
		} else {
			return true
		}
	}

	return false
}
