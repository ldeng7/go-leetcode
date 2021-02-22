func minimumJumps(forbidden []int, a int, b int, x int) (ans int) {
	n := x
	for _, f := range forbidden {
		if f > n {
			n = f
		}
	}
	n += b << 1
	v := make([][2]bool, n+1)
	for _, f := range forbidden {
		v[f][0], v[f][1] = true, true
	}
	v[0][0] = true

	o := 0
	q := make([]uint16, 1, 16)
	q[0] = 0
	for 0 != len(q) {
		for l := len(q); l > 0; l-- {
			t := q[0]
			q = q[1:]
			i, d := int(t>>1), t&1
			if i == x {
				return o
			}
			if j := i + a; j <= n && !v[j][0] {
				q = append(q, uint16(j<<1))
				v[j][0] = true
			}
			if d == 0 {
				if j := i - b; j >= 0 && !v[j][1] {
					q = append(q, uint16(j<<1)|1)
					v[j][1] = true
				}
			}
		}
		o++
	}
	return -1
}
