import "math"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func minDays(bloomDay []int, m int, k int) int {
	n := len(bloomDay)
	if m*k > n {
		return -1
	}
	l, r := math.MaxInt64, math.MinInt64
	for _, d := range bloomDay {
		l, r = min(l, d), max(r, d)
	}
	r++
	for l < r {
		mi := l + (r-l)>>1
		d := 0
		for i, j := 0, 0; i <= n; i++ {
			if i == n || bloomDay[i] > mi {
				j, d = i+1, d+(i-j)/k
			}
		}
		if d >= m {
			r = mi
		} else {
			l = mi + 1
		}
	}
	return l
}
