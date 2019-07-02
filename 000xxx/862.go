import "math"

func shortestSubarray(A []int, K int) int {
	l := len(A)
	if 0 == l {
		return -1
	}
	o, b, e := math.MaxInt64, 0, 0
	s, q := make([]int, l+1), make([]int, l+1)
	for i, a := range A {
		s[i+1] = s[i] + a
	}
	for i := 0; i <= l; i++ {
		if b != e {
			for b != e && s[i] <= s[q[e-1]] {
				e--
			}
			for b != e && s[i]-s[q[b]] >= K {
				if t := i - q[b]; t < o {
					o = t
				}
				b++
			}
		}
		q[e], e = i, e+1
	}
	if o == math.MaxInt64 {
		return -1
	}
	return o
}
