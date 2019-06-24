func magicalString(n int) int {
	if n <= 0 {
		return 0
	} else if n <= 3 {
		return 1
	}
	o, h, t, m := 1, 2, 3, 1
	ar := []int{1, 2, 2}
	for t < n {
		for i := 0; i < ar[h]; i++ {
			ar = append(ar, m)
			if m == 1 && t < n {
				o++
			}
			t++
		}
		m, h = m^3, h+1
	}
	return o
}
