import "math"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func mctFromLeafValues(arr []int) int {
	o, s := 0, []int{math.MaxInt64}
	for _, a := range arr {
		for s[len(s)-1] <= a {
			e := len(s) - 1
			b := s[e]
			s = s[:e]
			o += b * min(s[e-1], a)
		}
		s = append(s, a)
	}
	for i := 2; i < len(s); i++ {
		o += s[i] * s[i-1]
	}
	return o
}
