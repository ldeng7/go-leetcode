import "sort"

func sumSubseqWidths(A []int) int {
	sort.Ints(A)
	o, s, p := 0, 0, 1
	for i := len(A) - 2; i >= 0; {
		for j := 0; j < 22 && i >= 0; j, i = j+1, i-1 {
			p <<= 1
			s = (s << 1) + (p-1)*(A[i+1]-A[i])
			o += s
		}
		p, s, o = p%1000000007, s%1000000007, o%1000000007
	}
	return o
}
