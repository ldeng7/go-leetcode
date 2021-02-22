func maximumTime(time string) string {
	a, b, c, d := time[0], time[1], time[3], time[4]
	if a == '?' {
		if b == '?' || b <= '3' {
			a = '2'
		} else {
			a = '1'
		}
	}
	if b == '?' {
		if a == '2' {
			b = '3'
		} else {
			b = '9'
		}
	}
	if c == '?' {
		c = '5'
	}
	if d == '?' {
		d = '9'
	}
	return string([]byte{a, b, ':', c, d})
}
