func kthSmallest(mat [][]int, k int) int {
	m := len(mat)
	var cal func(int, int, int, int) int
	cal = func(mid, i, s, k int) int {
		if s > mid {
			return 0
		} else if i == m {
			return 1
		}
		o, ar := 0, mat[i]
		for _, n := range ar {
			t := cal(mid, i+1, s+n, k-o)
			if t == 0 {
				break
			}
			o += t
			if o > k {
				break
			}
		}
		return o
	}

	o, l, r := -1, m, m*5000
	for l <= r {
		mid := l + (r-l)>>1
		if cal(mid, 0, 0, k) >= k {
			o, r = mid, mid-1
		} else {
			l = mid + 1
		}
	}
	return o
}
