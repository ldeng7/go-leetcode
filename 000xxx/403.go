func canCross(stones []int) bool {
	l, k := len(stones), 0
	m := map[int]map[int]bool{}
	t := make([]int, l)
	m[0] = map[int]bool{0: true}
	for i := 1; i < l; i++ {
		m[i] = map[int]bool{}
		si := stones[i]
		for t[k]+1 < si-stones[k] {
			k++
		}
		for j := k; j < i; j++ {
			d := si - stones[j]
			if m[j][d-1] || m[j][d] || m[j][d+1] {
				m[i][d] = true
				if d > t[i] {
					t[i] = d
				}
			}
		}
	}
	return t[l-1] > 0
}
