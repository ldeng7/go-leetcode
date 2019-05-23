func prefixesDivBy5(A []int) []bool {
	out := make([]bool, len(A))
	n := 0
	for i, a := range A {
		n = (n<<1 + a) % 5
		out[i] = n == 0
	}
	return out
}
