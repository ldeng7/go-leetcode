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

func intervalIntersection(A [][]int, B [][]int) [][]int {
	o, la, lb := [][]int{}, len(A), len(B)
	for i, j := 0, 0; i < la && j < lb; {
		a, b := A[i], B[j]
		mi, ma := max(a[0], b[0]), min(a[1], b[1])
		if ma >= mi {
			o = append(o, []int{mi, ma})
		}
		if a[1] > b[1] {
			j += 1
		} else if a[1] < b[1] {
			i += 1
		} else {
			i, j = i+1, j+1
		}
	}
	return o
}
