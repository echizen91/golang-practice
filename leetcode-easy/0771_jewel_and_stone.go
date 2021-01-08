package easy

import "strings"

func numJewelsInStones(jewels string, stones string) int {
	var result int
	for _, c := range stones {
		if strings.Contains(jewels, string(c)) {
			result++
		}
	}
	return result
}
