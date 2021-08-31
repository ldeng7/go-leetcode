func sumInner(v, i int) int {
	if i >= v-1 {
		return (v-1)*v/2 + i + 1 - v
	}
	return i * (v - i + v - 1) / 2
}

func sum(v, n, i int) int {
	return sumInner(v, i) + v + sumInner(v, n-i-1)
}

func cal(n, i, ms, l, r int) int {
	if l == r {
		return l
	} else if r-l == 1 {
		if sum(r, n, i) <= ms {
			return r
		} else if sum(l, n, i) <= ms {
			return l
		}
		return 0
	}
	m := l + (r-l)>>1
	if sum(m, n, i) > ms {
		return cal(n, i, ms, l, m)
	}
	return cal(n, i, ms, m, r)
}

func maxValue(n int, index int, maxSum int) int {
	return cal(n, index, maxSum, 1, maxSum)
}
