package medium

/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0).
Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.

Notice that you may not slant the container.

Solution:
Runtime: 16 ms, faster than 56.79% of Go online submissions for Container With Most Water.
Memory Usage: 5.9 MB, less than 100.00% of Go online submissions for Container With Most Water.

Can I do better?
*/

func minCalc(x int, y int) int {
	if x <= y {
		return x
	}
	return y
}

func maxArea(height []int) int {
	var l, max int
	r := len(height) - 1
	for l < r {
		area := minCalc(height[l], height[r]) * (r - l)
		if max < area {
			max = area
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}
