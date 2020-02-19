func maximum69Number(num int) int {
	o, b := 0, true
	for num != 0 {
		var c int
		if num >= 1000 {
			c = 1000
		} else if num >= 100 {
			c = 100
		} else if num >= 10 {
			c = 10
		} else {
			c = 1
		}
		d := num / c
		num %= c
		if d == 6 && b {
			d, b = 9, false
		}
		o = o*10 + d
	}
	return o
}
