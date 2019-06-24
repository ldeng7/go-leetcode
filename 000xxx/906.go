import "strconv"

func reverse(n int) int {
	o := 0
	for ; n > 0; n /= 10 {
		o = o*10 + n%10
	}
	return o
}

func tenPowerFloor(n int) int {
	if n < 10 {
		return 1
	} else if n < 100 {
		return 10
	} else if n < 1000 {
		return 100
	} else if n < 10000 {
		return 1000
	} else {
		return 10000
	}
}

func superpalindromesInRange(L string, R string) int {
	l, _ := strconv.Atoi(L)
	r, _ := strconv.Atoi(R)
	o := 0
	for i := 0; i < 100000; i++ {
		f := tenPowerFloor(i)
		var a, b int
		if i != 1 {
			a, b = i*f+reverse(i/10), i*f*10+reverse(i)
		} else {
			a, b = i, i*10+i
		}
		a, b = a*a, b*b
		if a > r {
			break
		}
		if a >= l && a == reverse(a) {
			o += 1
		}
		if b < r && b >= l && b == reverse(b) {
			o += 1
		}
	}
	return o
}
