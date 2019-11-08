import "math"

func treeDiameter(edges [][]int) int {
	m := map[int][]int{}
	for _, e := range edges {
		a, b := e[0], e[1]
		m[a] = append(m[a], b)
		m[b] = append(m[b], a)
	}
	o := 0

	var cal func(int, int) int
	cal = func(n, p int) int {
		ar := m[n]
		if 1 == len(ar) && p == ar[0] {
			return 1
		}
		ma, ma1 := 0, 0
		for _, i := range ar {
			if i == p {
				continue
			}
			j := cal(i, n)
			if j > ma {
				ma, ma1 = j, ma
			} else if j > ma1 {
				ma1 = j
			}
			if t := ma + ma1; t > o {
				o = t
			}
		}
		return ma + 1
	}

	cal(edges[0][0], math.MaxInt64)
	return o
}
