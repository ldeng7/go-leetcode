func stoneGameVII(stones []int) int {
	n := len(stones)
	ds, ss := make([]int, n), make([]int, n)
	copy(ss, stones)
	for d, b := 1, n&1 == 1; d < n; d, b = d+1, !b {
		for i := 0; i < n-d; i++ {
			l, r := ds[i+1], ds[i]
			if b {
				l, r = l-ss[i+1], r-ss[i]
			} else {
				l, r = l+ss[i+1], r+ss[i]
			}
			if (b && l <= r) || ((!b) && l > r) {
				ds[i] = l
			} else {
				ds[i] = r
			}
			ss[i] += stones[i+d]
		}
	}
	return ds[0]
}
