func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := map[string]map[string]float64{}
	o := []float64{}
	for i, equ := range equations {
		if ms, ok := m[equ[0]]; ok {
			ms[equ[1]] = values[i]
		} else {
			m[equ[0]] = map[string]float64{equ[1]: values[i]}
		}
		if ms, ok := m[equ[1]]; ok {
			ms[equ[0]] = 1 / values[i]
		} else {
			m[equ[1]] = map[string]float64{equ[0]: 1 / values[i]}
		}
	}
	for _, q := range queries {
		s := map[string]bool{}
		var cal func(string, string) float64
		cal = func(u, d string) float64 {
			mu, ok := m[u]
			if !ok {
				return -1
			}
			if v, ok := mu[d]; ok {
				return v
			}
			for e, v := range mu {
				if s[e] {
					continue
				}
				s[e] = true
				t := cal(e, d)
				if t > 0 {
					return t * v
				}
			}
			return -1
		}
		t := cal(q[0], q[1])
		if t <= 0 {
			t = -1
		}
		o = append(o, t)
	}
	return o
}
