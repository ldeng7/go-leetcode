func dailyTemperatures(ts []int) []int {
	out := make([]int, len(ts))
	s := []int{}
	for i := 0; i < len(ts); i++ {
		for len(s) > 0 && ts[i] > ts[s[len(s)-1]] {
			t := s[len(s)-1]
			s = s[:len(s)-1]
			out[t] = i - t
		}
		s = append(s, i)
	}
	return out
}
