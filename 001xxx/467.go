func getProbability(balls []int) float64 {
	l := (len(balls) << 1) + 1
	n := 0
	for _, b := range balls {
		n += b
	}
	n &^= 1
	facts := make([]float64, n+1)
	facts[0] = 1.0
	for i := 1; i <= n; i++ {
		facts[i] = facts[i-1] * float64(i)
	}

	t := make([][]float64, n+1)
	for i := 0; i <= n; i++ {
		t[i] = make([]float64, l)
	}
	t[0][len(balls)] = 1.0
	s := 0
	for _, b := range balls {
		tn := make([][]float64, n+1)
		for i := 0; i <= n; i++ {
			tn[i] = make([]float64, l)
		}

		for i := 0; i <= b; i++ {
			d := 0
			if i == 0 {
				d = -1
			} else if i == b {
				d = 1
			}

			for j := 0; j <= n; j++ {
				for k := 0; k < l; k++ {
					e := t[j][k]
					if e == 0 {
						continue
					}
					tn[j+i][k+d] += e * facts[j+i] * facts[s-j+b-i] /
						(facts[j] * facts[i] * facts[s-j] * facts[b-i])
				}
			}
		}
		t = tn
		s += b
	}

	f := facts[n]
	for _, b := range balls {
		f /= facts[b]
	}
	return t[n>>1][len(balls)] / f
}
