func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minDominoRotations(A []int, B []int) int {
	l, as, bs, ns := len(A), [7]int{}, [7]int{}, [7]int{}
	for i, a := range A {
		b := B[i]
		if a == b {
			ns[a]++
		} else {
			as[a]++
			bs[b]++
			ns[a]++
			ns[b]++
		}
	}
	for n, c := range ns {
		if c == l {
			return min(as[n], bs[n])
		}
	}
	return -1
}
