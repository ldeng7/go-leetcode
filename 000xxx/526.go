func countArrangement(N int) int {
	ar := make([]int, N)
	for i := 0; i < N; i++ {
		ar[i] = i + 1
	}
	var cal func(n int) int
	cal = func(n int) int {
		if n <= 0 {
			return 1
		}
		o := 0
		for i := 0; i < n; i++ {
			if n%ar[i] == 0 || ar[i]%n == 0 {
				ar[i], ar[n-1] = ar[n-1], ar[i]
				o += cal(n - 1)
				ar[i], ar[n-1] = ar[n-1], ar[i]
			}
		}
		return o
	}
	return cal(N)
}
