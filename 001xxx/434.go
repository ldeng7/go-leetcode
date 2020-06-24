func numberWays(hats [][]int) int {
	t := [41][1024]int{}
	ar := [41][]int{}
	n := 1 << len(hats)
	for i, hs := range hats {
		for _, h := range hs {
			ar[h] = append(ar[h], i)
		}
	}
	t[0][0] = 1
	for i := 1; i <= 40; i++ {
		a, ap := &(t[i]), &(t[i-1])
		copy((*a)[:], (*ap)[:])
		for _, m := range ar[i] {
			m = 1 << m
			for j := 0; j < n; j++ {
				if (j & m) == 0 {
					(*a)[j|m] = ((*a)[j|m] + (*ap)[j]) % 1000000007
				}
			}
		}
	}
	return t[40][n-1]
}
