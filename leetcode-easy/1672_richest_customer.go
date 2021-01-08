package easy

func maximumWealth(accounts [][]int) int {
	var highest int
	for _, account := range accounts {
		var total int
		for _, m := range account {
			total += m
		}
		if highest < total {
			highest = total
		}
	}
	return highest
}
