package medium

import "math"

/*
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Solution:
Runtime: 8 ms, faster than 91.29% of Go online submissions for Minimum Path Sum.
Memory Usage: 3.9 MB, less than 5.30% of Go online submissions for Minimum Path Sum.
*/

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
				continue
			}
			if j == 0 {
				grid[i][j] += grid[i-1][j]
				continue
			}
			top := grid[i][j] + grid[i-1][j]
			left := grid[i][j] + grid[i][j-1]
			grid[i][j] = int(math.Min(float64(top), float64(left)))
		}
	}
	return grid[m-1][n-1]
}
