func minSteps(n int) int {
	o := 0
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			o += i
			n /= i
		}
	}
	return o
}
