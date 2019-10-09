func cnt(i uint32) uint32 {
	i = (i & 0x55555555) + ((i >> 1) & 0x55555555)
	i = (i & 0x33333333) + ((i >> 2) & 0x33333333)
	i = (i & 0x0f0f0f0f) + ((i >> 4) & 0x0f0f0f0f)
	i = (i & 0x00ff00ff) + ((i >> 8) & 0x00ff00ff)
	i = (i & 0x0000ffff) + ((i >> 16) & 0x0000ffff)
	return i
}

var ms = [100001]uint32{}

func canMakePaliQueries(s string, queries [][]int) []bool {
	o, l := make([]bool, len(queries)), len(s)
	var m uint32
	for i := 0; i < l; i++ {
		m ^= 1 << (s[i] - 'a')
		ms[i+1] = m
	}
	for i, q := range queries {
		o[i] = cnt(ms[q[0]]^ms[q[1]+1])>>1 <= uint32(q[2])
	}
	return o
}
