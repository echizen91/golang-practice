package easy

/*
Suppose you have a long flowerbed in which some of the plots are planted and some are not. However, flowers cannot be planted in adjacent plots - they would compete for water and both would die.

Given a flowerbed (represented as an array containing 0 and 1, where 0 means empty and 1 means not empty), and a number n, return if n new flowers can be planted in it without violating the no-adjacent-flowers rule.

Note:
The input array won't violate no-adjacent-flowers rule.
The input array size is in the range of [1, 20000].
n is a non-negative integer which won't exceed the input array size.
*/

/*
Runtime: 16 ms, faster than 85.51% of Go online submissions for Can Place Flowers.
Memory Usage: 6.1 MB, less than 100.00% of Go online submissions for Can Place Flowers.
*/
func canPlaceFlowers(flowerbed []int, n int) bool {
	// Base case
	if len(flowerbed) == 1 {
		if flowerbed[0] == 0 {
			return true
		}
		if n == 0 {
			return true
		}
		return false
	}

	for i := 0; i < len(flowerbed); i++ {
		// all done
		if n < 1 {
			return true
		}

		// can plant one flower if first 2 are empty
		if i == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
			continue
		} else if i == 0 {
			continue
		}
		// can plant one flower if last 2 are empty
		if i == len(flowerbed)-1 && flowerbed[i-1] == 0 && flowerbed[i] == 0 {
			n--
			continue
		} else if i == len(flowerbed)-1 {
			continue
		}

		if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 && flowerbed[i] == 0 {
			flowerbed[i] = 1
			n--
			continue
		}
	}

	if n < 1 {
		return true
	}

	return false
}
