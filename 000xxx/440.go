func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func findKthNumber(n int, k int) int {
	i := 1
	k--
	for k > 0 {
		d, b, e := 0, i, i+1
		for b <= n {
			d += min(n+1, e) - b
			b, e = b*10, e*10
		}
		if d <= k {
			i, k = i+1, k-d
		} else {
			i, k = i*10, k-1
		}
	}
	return i
}
