func average(salary []int) float64 {
	l := len(salary)
	s := salary[0]
	mi, ma := s, s
	for i := 1; i < l; i++ {
		a := salary[i]
		if a < mi {
			mi = a
		} else if a > ma {
			ma = a
		}
		s += a
	}
	return float64(s-mi-ma) / float64(l-2)
}
