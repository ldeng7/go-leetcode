func subarraysDivByK(A []int, K int) int {
	l := len(A)
	ss := make([]int, l+1)
	for i, a := range A {
		ss[i+1] = ss[i] + a
	}
	cs := make([]int, K)
	for _, s := range ss {
		cs[(s%K+K)%K]++
	}
	o := 0
	for _, c := range cs {
		o += c * (c - 1) / 2
	}
	return o
}
