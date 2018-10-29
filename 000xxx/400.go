import "strconv"

func findNthDigit(n int) int {
	l, c, i := 1, 9, 1
	for n > l*c {
		n -= l * c
		l++
		c *= 10
		i *= 10
	}
	i += (n - 1) / l
	s := []byte(strconv.Itoa(i))
	return int(s[(n-1)%l] - '0')
}
