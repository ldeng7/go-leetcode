import "math/bits"

func g(n, t int) int {
	if t == 0 {
		return 1 - n
	} else if m := 1 << t; n&m != 0 {
		return minimumOneBitOperations(n ^ m)
	} else {
		return m + g(n, t-1)
	}
}

func minimumOneBitOperations(n int) int {
	if n <= 1 {
		return n
	}
	t := 32 - bits.LeadingZeros32(uint32(n)) - 1
	m := 1 << t
	return m + g(n^m, t-1)
}
