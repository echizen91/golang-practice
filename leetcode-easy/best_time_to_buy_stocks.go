package easy

/*
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.
*/

/*
Runtime: 212 ms, faster than 7.45% of Go online submissions for Best Time to Buy and Sell Stock.
Memory Usage: 3.1 MB, less than 5.64% of Go online submissions for Best Time to Buy and Sell Stock.
*/
func slow(prices []int) int {
	var global int
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if global > (prices[i] - prices[j]) {
				global = prices[i] - prices[j]
			}
		}
	}
	if global < 0 {
		global *= -1
	}
	return global
}

/*
Runtime: 4 ms, faster than 94.57% of Go online submissions for Best Time to Buy and Sell Stock.
Memory Usage: 3.1 MB, less than 5.53% of Go online submissions for Best Time to Buy and Sell Stock.
*/
func fastProfit(prices []int) int {
	var global int
	current := -1
	for i, p := range prices {
		if current < 0 {
			current = p
			continue
		}
		temp := p - (prices[i-1])

		if temp < 1 {
			continue
		}

		if current > prices[i-1] {
			current = prices[i-1]
			if global < temp {
				global = temp
			}
		} else {
			global = max(p-current, global)
		}
	}
	return global
}

func maxProfit(prices []int) int {
	return fastProfit(prices)
}
