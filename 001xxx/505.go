func minInteger(num string, k int) string {
	l := len(num)
	t := [10][]int{}
	for i := 0; i < 10; i++ {
		t[i] = make([]int, 0, 8)
	}
	pos, ofs := [10]int{}, [10]int{}
	prev := make([]int, l)
	bs := make([]byte, 0, l)
	for i := 0; i < l; i++ {
		c := num[i] - '0'
		t[c] = append(t[c], i)
	}

	for 0 != k && len(bs) < l {
		for i := byte(0); i < 10; i++ {
			ti, pi := t[i], pos[i]
			if pi >= len(ti) || ti[pi]+ofs[i] > len(bs)+k {
				continue
			}
			tp := ti[pi]
			k -= tp + ofs[i] - len(bs)
			bs = append(bs, i+'0')
			prev[tp]--
			for j := byte(0); j < 10; j++ {
				if pos[j] < len(t[j]) && tp >= t[j][pos[j]] {
					ofs[j]++
				}
			}

			pi, pos[i] = pi+1, pi+1
			if pi < len(ti) {
				je := ti[pi]
				for j := tp; j < je; j++ {
					ofs[i] += prev[j]
				}
			}
			break
		}
	}

	for i := 0; i < l; i++ {
		if 0 == prev[i] {
			bs = append(bs, num[i])
		}
	}
	return string(bs)
}
