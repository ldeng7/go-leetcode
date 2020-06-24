func cal(n, start int) int {
	if n&1 == 0 {
		return (n >> 1) & 1
	}
	return ((n >> 1) & 1) ^ (start + n - 1)
}

func xorOperation(n int, start int) int {
	s := start >> 1
	var o int
	if s&1 == 0 {
		o = cal(n, s) << 1
	} else {
		o = (cal(n+1, s-1) ^ (s - 1)) << 1
	}
	if n&start&1 == 1 {
		o++
	}
	return o
}
