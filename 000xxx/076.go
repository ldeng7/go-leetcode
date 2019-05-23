import "math"

func minWindow(s string, t string) string {
	var out string
	mt := [128]int{}
	l, c, min := 0, 0, math.MaxInt64
	for _, b := range t {
		mt[b]++
	}
	for i, b := range s {
		mt[b]--
		if mt[b] >= 0 {
			c++
		}
		for c == len(t) {
			if min > i-l+1 {
				min = i - l + 1
				out = s[l : l+min]
			}
			bl := s[l]
			mt[bl]++
			if mt[bl] > 0 {
				c--
			}
			l++
		}
	}
	return out
}
