import "sort"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}

func minCost(n int, cuts []int) int {
	sort.Ints(cuts)
	l := len(cuts)
	ar := make([]int, l+1)
	ar[0] = cuts[0]
	for i := 1; i < l; i++ {
		ar[i] = cuts[i] - cuts[i-1]
	}
	ar[l] = n - cuts[l-1]

	l++
	ps := make([]int, l+1)
	for i := 1; i <= l; i++ {
		ps[i] = ps[i-1] + ar[i-1]
	}
	t, u := make([][]int, l), make([][]int, l)
	for i := 0; i < l; i++ {
		t[i], u[i] = make([]int, l), make([]int, l)
		u[i][i] = i
	}
	for i := 2; i <= l; i++ {
		for j, k := 0, i-1; k < l; j, k = j+1, k+1 {
			t[j][k] = 1e9
			for h, he := max(j, u[j][k-1]), min(k-1, u[j+1][k]); h <= he; h++ {
				if v := t[j][h] + t[h+1][k] + ps[k+1] - ps[j]; v < t[j][k] {
					t[j][k], u[j][k] = v, h
				}
			}
		}
	}
	return t[0][l-1]
}
