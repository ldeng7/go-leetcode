import "strconv"

func numDecodings(s string) int {
	i, _ := strconv.Atoi(s)
	if 0 == i || s[0] == '0' {
		return 0
	}
	a, b := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			a = 0
		}
		if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
			a, b = a+b, a
		} else {
			b = a
		}
	}
	return a
}
