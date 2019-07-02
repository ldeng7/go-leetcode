func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func cmp(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

func maxTurbulenceSize(A []int) int {
	o, l := 1, len(A)
	for i, j := 1, 0; i < l; i++ {
		c := cmp(A[i-1], A[i])
		if i == l-1 || c*cmp(A[i], A[i+1]) != -1 {
			if c != 0 {
				o = max(o, i-j+1)
			}
			j = i
		}
	}
	return o
}
