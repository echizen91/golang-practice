package hard

/*
Given an unsorted integer array nums, find the smallest missing positive integer.

Follow up: Could you implement an algorithm that runs in O(n) time and uses constant extra space.?
*/

func firstMissingPositive(nums []int) int {
	seedNumber := 1
	numMap := make(map[int]int)
	for _, n := range nums {
		if n == seedNumber {
			success := false
			for !success {
				if _, ok := numMap[seedNumber+1]; ok {
					seedNumber++
					continue
				} else {
					seedNumber++
					success = true
				}
			}
			continue
		}
		numMap[n] = n
	}
	return seedNumber
}
