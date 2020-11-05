package easy

/*
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed,
the only constraint stopping you from robbing each of them is that adjacent houses
have security system connected and it will automatically contact the police if two
adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house,
determine the maximum amount of money you can rob tonight without alerting the police.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber.
Memory Usage: 2 MB, less than 100.00% of Go online submissions for House Robber.
*/
func rob(nums []int) int {
	var global, previous int
	for _, n := range nums {
		temp := global
		global = max(previous+n, global)
		previous = temp
	}
	return global
}
