func findMinFibonacciNumbers(k int) int {
	a, b := 0, 1
	for b <= k {
		a, b = b, a
		b += a
	}
	if k == a {
		return 1
	}
	return findMinFibonacciNumbers(k-a) + 1
}
