import "math"

func flipgame(fronts []int, backs []int) int {
	o := math.MaxInt64
	s := map[int]bool{}
	for i, f := range fronts {
		if f == backs[i] {
			s[f] = true
		}
	}
	for _, f := range fronts {
		if _, ok := s[f]; !ok && f < o {
			o = f
		}
	}
	for _, b := range backs {
		if _, ok := s[b]; !ok && b < o {
			o = b
		}
	}
	if o == math.MaxInt64 {
		return 0
	}
	return o
}
