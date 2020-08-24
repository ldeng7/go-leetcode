import "math"

func enc(i int, p float64) uint64 {
	return (uint64(i) << 32) | uint64(math.Float32bits(float32(p)))
}

func dec(k uint64) (int, float64) {
	return int(k >> 32), float64(math.Float32frombits(uint32(k & 0xffffffff)))
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	m := make([][]uint64, n)
	for i, e := range edges {
		a, b, p := e[0], e[1], succProb[i]
		m[a] = append(m[a], enc(b, p))
		m[b] = append(m[b], enc(a, p))
	}
	ar, bs, q := make([]float64, n), make([]bool, n), make([]int, 1, 64)
	ar[start], bs[start], q[0] = 1, true, start
	for 0 != len(q) {
		i := q[0]
		q = q[1:]
		for _, k := range m[i] {
			j, p := dec(k)
			if t := ar[i] * p; t > ar[j] {
				ar[j] = t
				if !bs[j] {
					bs[j] = true
					q = append(q, j)
				}
			}
		}
		bs[i] = false
	}
	return ar[end]
}
