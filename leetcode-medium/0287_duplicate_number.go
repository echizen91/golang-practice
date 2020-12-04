package medium

func findDuplicate(nums []int) int {
	seenMap := make(map[int]int, 0)
	for _, n := range nums {
		if _, ok := seenMap[n]; ok {
			return n
		}
		seenMap[n] = n
	}
	return 0
}
