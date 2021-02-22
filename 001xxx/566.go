func containsPattern(arr []int, m int, k int) bool {
	cal := func(l, r, i int) bool {
		for ; l < r; l, i = l+1, i+1 {
			if arr[l] != arr[i] {
				return false
			}
		}
		return true
	}

	l := len(arr) - m
	for i := 0; i < l; i++ {
		n := 1
		for j := i + m; j <= l; j += m {
			if cal(i, i+m, j) {
				n++
			} else {
				break
			}
			if n == k {
				return true
			}
		}
	}
	return false
}
