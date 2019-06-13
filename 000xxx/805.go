import "sort"

func splitArraySameAverage(A []int) bool {
	l, s, b := len(A), 0, false
	for _, a := range A {
		s += a
	}
	for i := 1; i <= l>>1 && !b; i++ {
		if s*i%l == 0 {
			b = true
		}
	}
	if !b {
		return false
	}

	sort.Ints(A)

	var cal func(s, n, b int) bool
	cal = func(s, n, b int) bool {
		if n == 0 {
			return s == 0
		}
		if A[b] > s/n {
			return false
		}
		for i := b; i < l-n+1; i++ {
			if i > b && A[i] == A[i-1] {
				continue
			}
			if cal(s-A[i], n-1, i+1) {
				return true
			}
		}
		return false
	}

	for i := 1; i <= l>>1; i++ {
		if s*i%l == 0 && cal(s*i/l, i, 0) {
			return true
		}
	}
	return false
}
