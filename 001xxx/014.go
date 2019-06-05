func maxScoreSightseeingPair(A []int) int {
	o, x := 0, A[0]
	for j := 1; j < len(A); j++ {
		a := A[j]
		ot := x + a - j
		if ot > o {
			o = ot
		}
		x1 := a + j
		if x1 > x {
			x = x1
		}
	}
	return o
}
