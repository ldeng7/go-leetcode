func C(m, n int) int {
	if n-m < m {
		m = n - m
	}
	o := 1
	for i := n; i > n-m; i-- {
		o *= i
	}
	d := 1
	for i := m; i > 1; i-- {
		d *= i
	}
	return o / d
}

func paintingPlan(n int, k int) int {
	if k == 0 || k == n*n {
		return 1
	} else if k < n {
		return 0
	}
	o := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if n*(i+j)-i*j == k {
				o += C(i, n) * C(j, n)
			}
		}
	}
	return o
}
