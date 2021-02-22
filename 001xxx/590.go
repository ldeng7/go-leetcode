func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minSubarray(nums []int, p int) int {
	s := 0
	for _, n := range nums {
		s += n
	}
	s1 := s % p
	if s1 == 0 {
		return 0
	}
	m := map[int]int{0: -1}
	l := len(nums)
	o, s := l, 0
	for i, n := range nums {
		s = (s + n) % p
		k := (s - s1 + p) % p
		if v, ok := m[k]; ok {
			o = min(o, i-v)
		}
		m[s] = i
	}
	if o >= l {
		return -1
	}
	return o
}
