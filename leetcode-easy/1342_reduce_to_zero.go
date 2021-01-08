package easy

func numberOfSteps(num int) int {
	var counter int
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		counter++
	}
	return counter
}
