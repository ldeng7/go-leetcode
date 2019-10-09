func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start > destination {
		start, destination = destination, start
	}
	s, s1 := 0, 0
	for i, d := range distance {
		s += d
		if i >= start && i < destination {
			s1 += d
		}
	}
	if s0 := s - s1; s0 < s1 {
		s1 = s0
	}
	return s1
}
