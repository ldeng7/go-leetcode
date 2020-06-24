func min(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func enc(s, e int) uint64 {
	return (uint64(s) << 32) | uint64(e)
}
func dec(k uint64) (int, int) {
	return int(k >> 32), int(k & 0xffffffff)
}

type MyCalendarTwo struct {
	calendar, overlap []uint64
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{make([]uint64, 0, 64), make([]uint64, 0, 64)}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for _, k := range this.overlap {
		if s, e := dec(k); start < e && end > s {
			return false
		}
	}
	for _, k := range this.calendar {
		if s, e := dec(k); start < e && end > s {
			this.overlap = append(this.overlap, enc(max(start, s), min(end, e)))
		}
	}
	this.calendar = append(this.calendar, enc(start, end))
	return true
}
