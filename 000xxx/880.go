func decodeAtIndex(S string, K int) string {
	i, c := 0, 0
	for ; c < K; i++ {
		b := S[i]
		if b >= '0' && b <= '9' {
			c *= int(b - '0')
		} else {
			c++
		}
	}
	for i > 0 {
		i--
		b := S[i]
		if b >= '0' && b <= '9' {
			c /= int(b - '0')
			K %= c
		} else {
			if K%c == 0 {
				return S[i : i+1]
			}
			c--
		}
	}
	return "oh no"
}
