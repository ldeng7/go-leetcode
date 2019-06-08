func superPow(a int, b []int) int {
	var pow func(x, y int) int
	pow = func(x, y int) int {
		switch y {
		case 0:
			return 1
		case 1:
			return x % 1337
		default:
			return pow(x%1337, y/2) * pow(x%1337, y-y/2) % 1337
		}
	}
	out := 1
	for _, e := range b {
		out = pow(out, 10) * pow(a, e) % 1337
	}
	return out
}
