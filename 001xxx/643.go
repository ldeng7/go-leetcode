func gcd(a, b int) int {
	if 0 == b {
		return a
	}
	return gcd(b, a%b)
}

func kthSmallestPath(destination []int, k int) string {
	v, h := destination[0], destination[1]
	n := v + h
	o := make([]byte, 0, n)
	for ; h != 0; n-- {
		n1, h1, s, r := n-1, h-1, 1, 1
		if n1-h1 < h1 {
			h1 = n1 - h1
		}
		if h1 != 0 {
			for ; h1 != 0; n1, h1 = n1-1, h1-1 {
				s, r = s*n1, r*h1
				m := gcd(s, r)
				s, r = s/m, r/m
			}
		}
		if k > s {
			o = append(o, 'V')
			k, v = k-s, v-1
		} else {
			o = append(o, 'H')
			h--
		}
	}
	for ; v != 0; v-- {
		o = append(o, 'V')
	}
	return string(o)
}
