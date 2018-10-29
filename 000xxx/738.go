import "strconv"

func monotoneIncreasingDigits(n int) int {
	bs := []byte(strconv.Itoa(n))
	j := len(bs)
	for i := len(bs) - 1; i > 0; i-- {
		if bs[i] >= bs[i-1] {
			continue
		}
		bs[i-1]--
		j = i
	}
	for i := j; i < len(bs); i++ {
		bs[i] = '9'
	}
	out, _ := strconv.Atoi(string(bs))
	return out
}
