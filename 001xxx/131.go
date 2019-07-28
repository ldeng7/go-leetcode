func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func cal(ar []int) int {
	ma, mi := ar[0], ar[0]
	for _, a := range ar {
		ma, mi = max(ma, a), min(mi, a)
	}
	return ma - mi
}

func maxAbsValExpr(arr1 []int, arr2 []int) int {
	l := len(arr1)
	s1, s2, d1, d2 := make([]int, l), make([]int, l), make([]int, l), make([]int, l)
	for i := 0; i < l; i++ {
		a1, a2 := arr1[i], arr2[i]
		s1[i], s2[i] = a1+a2+i, a1+a2-i
		d1[i], d2[i] = a1-a2+i, a1-a2-i
	}
	return max(max(cal(s1), cal(d1)), max(cal(s2), cal(d2)))
}
