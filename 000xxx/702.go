import "math"

func search(reader ArrayReader, target int) int {
	l, r := 0, math.MaxInt64
	for l < r {
		m := l + (r-l)>>1
		i := reader.get(m)
		if i == target {
			return m
		} else if i < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return -1
}
