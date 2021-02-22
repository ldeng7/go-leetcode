func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func eatenApples(apples []int, days []int) int {
	o, a1, d1 := 0, 0, 0
	for i, a := range apples {
		d := days[i]
		if a1 > 0 && a > 0 {
			o++
			if d1 >= d && a1 >= a {
				a1, d1 = a1+a-1, d1-1
			} else if d1 >= d && a1 <= a {
				a, d, d1 = a-1, d-1, d1-1
				a1 += min(a, d)
			} else if d1 <= d && a1 <= a {
				a1, d, d1 = a1+a-1, d-1, d-1
			} else {
				d, d1 = d-1, d-1
				a1 = a1 - 1 + min(a, d)
			}
		} else if a1 > 0 {
			o, a1, d1 = o+1, a1-1, d1-1
		} else if a > 0 {
			o, a1, d1 = o+1, a-1, d-1
		}
		if d1 <= 0 || a1 <= 0 {
			a1, d1 = 0, 0
		}
		if a1 > d1 {
			a1 = d1
		}
	}
	if a1 > 0 {
		o += a1
	}
	return o
}
