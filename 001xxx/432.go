func cal(n, x, y int) int {
	if n == 0 {
		return 0
	}
	d := n % 10
	if d == x {
		d = y
	}
	return cal(n/10, x, y)*10 + d
}

func maxDiff(num int) int {
	a, b := num, num
	ds := make([]int, 0, 9)
	for 0 != num {
		ds = append(ds, num%10)
		num /= 10
	}
	l := len(ds)
	for i := l - 1; i >= 0; i-- {
		if d := ds[i]; d != 9 {
			a = cal(a, d, 9)
			break
		}
	}
	if ds[l-1] != 1 {
		b = cal(b, ds[l-1], 1)
	} else {
		for i := l - 2; i >= 0; i-- {
			if d := ds[i]; d != 0 && d != 1 {
				b = cal(b, d, 0)
				break
			}
		}
	}
	return a - b
}
