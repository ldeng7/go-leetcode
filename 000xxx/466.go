func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	cs, ns := make([]int, n1+1), make([]int, n1+1)
	j, c := 0, 0
	for k := 1; k <= n1; k++ {
		for i := 0; i < len(s1); i++ {
			if s1[i] == s2[j] {
				j++
				if j == len(s2) {
					j, c = 0, c+1
				}
			}
		}
		cs[k], ns[k] = c, j
		for h := 0; h < k; h++ {
			if ns[h] == j {
				return ((cs[k]-cs[h])*((n1-h)/(k-h)) + cs[h+(n1-h)%(k-h)]) / n2
			}
		}
	}
	return cs[n1] / n2
}
