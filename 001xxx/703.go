func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minMoves(nums []int, k int) int {
	n := len(nums)
	ar := make([]int, 1, n+1)
	ar[0] = 0
	for i, num := range nums {
		if 0 != num {
			ar = append(ar, i)
		}
	}
	n = len(ar)
	t := make([]int, n)
	for i := 1; i < n; i++ {
		t[i] = ar[i] + t[i-1]
	}

	var o int = 1e10
	for l, r := 1, k; r < n; l, r = l+1, r+1 {
		m := l + (r-l)>>1
		d := m - l
		l1 := t[m-1] - t[l-1]
		s := d * (d + 1)
		if k&1 == 0 {
			l1 += ar[m]
			s += (d + 1)
		}
		o = min(o, t[r]-t[m]-l1-s)
	}
	return o
}
