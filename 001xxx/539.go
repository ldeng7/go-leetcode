func findKthPositive(arr []int, k int) int {
	if arr[0] > k {
		return k
	}
	n := len(arr)
	l, r := 0, n
	for l < r {
		m := l + (r-l)>>1
		if m >= n || arr[m]-m-1 >= k {
			r = m
		} else {
			l = m + 1
		}
	}
	return k + l
}
