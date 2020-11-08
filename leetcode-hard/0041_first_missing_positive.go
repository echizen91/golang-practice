package hard

/*
Given an unsorted integer array nums, find the smallest missing positive integer.

Follow up: Could you implement an algorithm that runs in O(n) time and uses constant extra space.?
*/

func firstMissingPositive(nums []int) int {
	// Set lowest positive number
	seedNumber := 1

	// Map to store positive numbers currently more than seed positive number
	numMap := make(map[int]int)

	// Ranging through nums
	for _, n := range nums {

		// if found a match from array to seedNumber
		if n == seedNumber {
			success := false

			// while
			for !success {
				// keep incrementing positive seed until no match is found in map
				if _, ok := numMap[seedNumber+1]; ok {
					seedNumber++
					continue
				} else {
					// break when no match is found
					seedNumber++
					success = true
				}
			}
			continue
		}
		// else insert number to map
		numMap[n] = n
	}

	// return the lowest possible seedNumber
	return seedNumber
}
