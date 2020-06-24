func countLargestGroup(n int) int {
	ar := [100]int{}
	for i := 1; i <= n; i++ {
		c, v := 0, i
		for v != 0 {
			c += v % 10
			v /= 10
		}
		ar[c]++
	}
	m := 0
	for _, a := range ar {
		if a > m {
			m = a
		}
	}
	o := 0
	for _, a := range ar {
		if a == m {
			o++
		}
	}
	return o
}
