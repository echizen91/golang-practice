package medium

/*
Initially on a notepad only one character 'A' is present. You can perform two operations on this notepad for each step:

Copy All: You can copy all the characters present on the notepad (partial copy is not allowed).
Paste: You can paste the characters which are copied last time.


Given a number n. You have to get exactly n 'A' on the notepad by performing the minimum number of steps permitted. Output the minimum number of steps to get n 'A'.
*/

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for 2 Keys Keyboard.
Memory Usage: 1.9 MB, less than 5.00% of Go online submissions for 2 Keys Keyboard.
*/
func minSteps(n int) int {
	// Base case
	if n == 1 {
		return 0
	}

	current, size, steps := 1, 1, 0
	for n > 0 {
		if steps == 0 {
			// Copy for base case
			steps++
			n--
			continue
		}
		if steps == 1 {
			// Paste for base case
			steps++
			current++
			n--
			continue
		}
		mod := n % current
		if mod == 0 {
			// Do copy & paste
			steps += 2
			size = current
			n -= size
			current *= 2
			continue
		}
		// else just paste
		steps++
		n -= size
		current += size
	}
	return steps
}
