package medium

/*
A robot is located at the top-left corner of a m x n grid (marked 'Start' in the diagram below).

The robot can only move either down or right at any point in time.
The robot is trying to reach the bottom-right corner of the grid (marked 'Finish' in the diagram below).

How many possible unique paths are there?

Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Paths.
Memory Usage: 1.9 MB, less than 7.29% of Go online submissions for Unique Paths.
*/

// Solution will exceed time limit (but workable)
func countPath(m int, n int, currentM int, currentN int) int {
	result := 0
	// Base cases
	if m == 1 && n == 1 {
		return 1
	}
	if m == 1 && n == 2 || m == 2 && n == 1 {
		return 1
	}

	if currentM == m && currentN == n {
		return 1
	}

	if currentM <= m {
		result += countPath(m, n, currentM+1, currentN)
	}
	if currentN <= n {
		result += countPath(m, n, currentM, currentN+1)
	}

	return result

}

func permutations(m int, n int) int {
	// Base conditions
	m, n = m-1, n-1
	if m == 0 || n == 0 {
		return 1
	}
	// ensure m is always greater,
	// otherwise some test cases result in overflow
	if m < n {
		m, n = n, m
	}

	d1, d2 := 1, 1
	// log.Println(m, n)
	for i := m + 1; i <= m+n; i++ {
		d1 *= i
	}
	for i := 1; i <= n; i++ {
		d2 *= i
	}
	return d1 / d2
}

func uniquePaths(m int, n int) int {
	// return countPath(m, n, 1, 1)
	return permutations(m, n)
}
