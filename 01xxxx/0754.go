func reachNumber(target int) int {
	t := target
	if t < 0 {
		t = -t
	}
	o, s := 0, 0
	for s < t || (s-t)&1 == 1 {
		o++
		s += o
	}
	return o
}
