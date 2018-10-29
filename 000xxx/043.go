func multiply(num1 string, num2 string) string {
	out := []byte{}
	bs1, bs2 := []byte(num1), []byte(num2)
	l1, l2 := len(num1), len(num2)
	k, c := l1+l2-2, 0
	sl := make([]int, l1+l2)
	for i, b1 := range bs1 {
		for j, b2 := range bs2 {
			sl[k-i-j] += int((b1 - '0') * (b2 - '0'))
		}
	}
	for i := 0; i < l1+l2; i++ {
		sl[i] += c
		c = sl[i] / 10
		sl[i] %= 10
	}
	i := l1 + l2 - 1
	for i >= 0 && sl[i] == 0 {
		i--
	}
	if i < 0 {
		return "0"
	}
	for i >= 0 {
		out = append(out, byte(sl[i])+'0')
		i--
	}
	return string(out)
}
