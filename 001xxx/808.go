const M = 1000000007

func power(b, p int) int {
	o := 1
	for ; p != 0; p >>= 1 {
		if p&1 == 1 {
			o = (o * b) % M
		}
		b = (b * b) % M
	}
	return o
}

func maxNiceDivisors(n int) int {
	if n > 3 {
		switch n % 3 {
		case 0:
			return power(3, n/3)
		case 1:
			return (4 * power(3, (n/3)-1)) % M
		case 2:
			return (2 * power(3, n/3)) % M
		}
	}
	return n
}
