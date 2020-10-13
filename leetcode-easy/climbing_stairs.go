package easy

/*
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Solution:
Runtime: 0 ms, faster than 100.00% of Go online submissions for Climbing Stairs.
Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Climbing Stairs.
*/

// ClimbingStairs : count the number of ways to climb to the top
func ClimbingStairs(n int) int {
	var fn, result int
	if n == 1 {
		return 1
	}
	fn = n - 2
	var prev, cur = 1, 1
	for i := 0; i <= fn; i++ {
		result = prev + cur
		prev, cur = cur, prev+cur
	}
	return result
}
