package easy

func subtractProductAndSum(n int) int {
	var sum int
	prod := 1
	for n > 0 {
		remainder := n % 10
		prod *= remainder
		sum += remainder
		n /= 10
	}
	return prod - sum
}
