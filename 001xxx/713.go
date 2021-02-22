func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minOperations(target []int, arr []int) int {
	m := map[int]int{}
	for i, t := range target {
		m[t] = i
	}
	t := make([]int, 0, 16)
	for _, a := range arr {
		v, ok := m[a]
		if !ok {
			continue
		} else if 0 == len(t) || v > t[len(t)-1] {
			t = append(t, v)
			continue
		}
		l, r := 0, len(t)-1
		k := r
		for l <= r {
			mi := l + (r-l)>>1
			if t[mi] >= v {
				k = min(k, mi)
				r = mi - 1
			} else {
				l = mi + 1
			}
		}
		t[k] = v
	}
	return len(target) - len(t)
}
