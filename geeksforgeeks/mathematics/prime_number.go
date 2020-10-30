package geekmath

/*
For a given number N check if it is prime or not. A prime number is a number which is only divisible by 1 and itself.
*/

func isPrime(n int) bool {
	if n == 1 || n == 4 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}

	for i := 5; i*i <= n; i += 6 {
		if n%5 == 0 {
			return false
		}
	}

	return true
}
