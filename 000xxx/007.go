func reverse(x int) int {
	m := false
	if x < 0 {
		m = true
		x = -x
	}
	ds := []int{}
	for x > 0 {
		ds = append(ds, x%10)
		x /= 10
	}
	o := 0
	b := 1
	for i := len(ds) - 1; i >= 0; i-- {
		o += ds[i] * b
		b *= 10
		if !m {
			if o >= 2147483648 {
				return 0
			}
		} else {
			if o > 2147483648 {
				return 0
			}
		}
	}
	if m {
		o = -o
	}
	return o
}