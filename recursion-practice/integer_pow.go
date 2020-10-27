package recursion

/*
raise a to the power of b recursively
*/

func recursivePow(a int, b int) int {
	if b < 1 {
		return 1
	}
	return a * recursivePow(a, b-1)
}
