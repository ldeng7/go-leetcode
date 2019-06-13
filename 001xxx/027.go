func longestArithSeqLength(A []int) int {
	o, l := 2, len(A)
	m := map[int][]int{}
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			d := A[j] - A[i]
			var v int
			if t, ok := m[d]; ok {
				if 0 != t[i] {
					t[j] = t[i] + 1
				} else {
					t[j] = 2
				}
				v = t[j]
			} else {
				t := make([]int, l)
				t[j], v, m[d] = 2, 2, t
			}
			if v > o {
				o = v
			}
		}
	}
	return o
}
