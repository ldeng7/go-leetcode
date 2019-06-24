import "sort"

func smallestRangeII(A []int, K int) int {
	if 1 == len(A) {
		return 0
	}
	min, max := A[0], A[0]
	for _, a := range A[1:] {
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
	}
	d := max - min
	if d >= K<<2 {
		return d - K<<1
	} else if d <= K {
		return d
	}

	t := []int{max - K<<1, min + K<<1}
	for _, a := range A {
		if (max-K<<1) < a && a < (min+K<<1) {
			t = append(t, a)
		}
	}
	sort.Ints(t)
	d = t[0] - t[1] + K<<1
	for i := 1; i < len(t)-1; i++ {
		d1 := t[i] - t[i+1] + K<<1
		if d1 < d {
			d = d1
		}
	}
	return d
}
