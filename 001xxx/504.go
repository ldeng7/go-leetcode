func numSubmat(mat [][]int) int {
	o, n := 0, len(mat[0])
	ar := make([]int, n)
	for _, line := range mat {
		for i, a := range line {
			if 0 != a {
				ar[i]++
			} else {
				ar[i] = 0
			}
		}
		t, st := make([]int, n), make([]int, 0, n)
		for i := 0; i < n; i++ {
			for 0 != len(st) && ar[st[len(st)-1]] >= ar[i] {
				st = st[:len(st)-1]
			}
			if 0 != len(st) {
				j := st[len(st)-1]
				t[i] = t[j] + ar[i]*(i-j)
			} else {
				t[i] = ar[i] * (i + 1)
			}
			st = append(st, i)
		}
		for _, a := range t {
			o += a
		}
	}
	return o
}
