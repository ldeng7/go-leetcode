func numSubarraysWithSum(A []int, S int) int {
	o, l := 0, len(A)
	ss, t := make([]int, l+1), make([]int, l+1)
	for i, a := range A {
		ss[i+1] = ss[i] + a
	}
	t[0] = 1
	for i := 1; i <= l; i++ {
		s := ss[i]
		if s >= S {
			o += t[s-S]
		}
		t[s]++
	}
	return o
}
