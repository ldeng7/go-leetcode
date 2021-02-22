func maxProfit(inventory []int, orders int) int {
	ma, mi := 0, 0
	for _, v := range inventory {
		if v > ma {
			ma = v
		}
	}
	var m int
	for {
		m = ma + (mi-ma)>>1
		s, t := 0, 0
		for _, v := range inventory {
			if v >= m {
				t++
				s += v - m
			}
		}
		if s > orders {
			mi = m + 1
		} else if s+t <= orders {
			ma = m - 1
		} else {
			break
		}
	}
	o, r := 0, 0
	for _, v := range inventory {
		if v > m {
			o = (o + (m+v+1)*(v-m)/2) % 1000000007
			r += v - m
		}
	}
	return (o + (orders-r)*m) % 1000000007
}
