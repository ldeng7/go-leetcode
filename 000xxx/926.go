func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minFlipsMonoIncr(S string) int {
	n0, n1, l := 0, 0, len(S)
	cs := make([]int, l)
	for i := l - 1; i >= 0; i-- {
		if S[i] == '0' {
			n0++
		} else {
			cs[i] = n0
		}
	}
	if n0 == len(S) || n0 == 0 {
		return 0
	}

	o := min(n0, l-n0)
	for i := 0; i < l; i++ {
		if S[i] == '0' {
			continue
		}
		o, n1 = min(o, n1+cs[i]), n1+1
	}
	return o
}
