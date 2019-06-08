func largestOverlap(A [][]int, B [][]int) int {
	o, l := 0, len(A)
	ara, arb, m := []int{}, []int{}, map[int]int{}
	for i := 0; i < l*l; i++ {
		if A[i/l][i%l] == 1 {
			ara = append(ara, i/l*100+i%l)
		}
		if B[i/l][i%l] == 1 {
			arb = append(arb, i/l*100+i%l)
		}
	}
	for _, a := range ara {
		for _, b := range arb {
			m[a-b]++
		}
	}
	for _, c := range m {
		if c > o {
			o = c
		}
	}
	return o
}
