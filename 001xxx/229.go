import "math"

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	ar := make([]int, 0, 64)
	for _, s1 := range slots1 {
		s1b, s1e := s1[0], s1[1]
		if s1e-s1b < duration {
			continue
		}
		for _, s2 := range slots2 {
			s2b, s2e := s2[0], s2[1]
			if b := s2b + duration; s1b < s2b && b <= s2e && b <= s1e {
				ar = append(ar, s2b)
				ar = append(ar, ar[0]+duration)
			} else if b := s1b + duration; s1b >= s2b && b <= s2e && b <= s1e {
				ar = append(ar, s1b)
				ar = append(ar, ar[0]+duration)
			}
		}
	}
	if len(ar) == 0 {
		return ar
	}
	m := math.MaxInt64
	for _, v := range ar {
		if m > v {
			m = v
		}
	}
	return []int{m, m + duration}
}
