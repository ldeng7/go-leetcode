type MyCalendarTwo struct {
	l1 map[complex128]bool
	l2 map[complex128]bool
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{map[complex128]bool{}, map[complex128]bool{}}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	s, e := float64(start), float64(end)
	for c, _ := range this.l2 {
		if s >= imag(c) || e <= real(c) {
			continue
		}
		return false
	}
	for c, _ := range this.l1 {
		if s >= imag(c) || e <= real(c) {
			continue
		}
		sn := s
		if real(c) > s {
			sn = real(c)
		}
		en := e
		if imag(c) < e {
			en = imag(c)
		}
		this.l2[complex(sn, en)] = true
	}
	this.l1[complex(s, e)] = true
	return true
}
