package easy

func xorOperation(n int, start int) int {
	if n == 1 {
		return start
	}

	result := start
	current := start
	for n > 1 {
		current += 2
		result ^= current
		n--
	}
	return result
}
