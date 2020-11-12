package medium

// easy solution: not optimized
func sortColors(nums []int) []int {
	var red, white, blue int
	for _, n := range nums {
		if n == 0 {
			red++
			continue
		}
		if n == 1 {
			white++
			continue
		}
		if n == 2 {
			blue++
			continue
		}
	}

	var index int

	for red > 0 || white > 0 || blue > 0 {
		if red > 0 {
			nums[index] = 0
			red--
			index++
			continue
		}
		if white > 0 {
			nums[index] = 1
			white--
			index++
			continue
		}
		if blue > 0 {
			nums[index] = 2
			blue--
			index++
			continue
		}
	}

	return nums
}
