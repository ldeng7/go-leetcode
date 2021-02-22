func canFormArray(arr []int, pieces [][]int) bool {
	m, n := len(arr), len(pieces)
	v := make([]bool, m)
	for i := 0; i < m; {
		j := 0
		for ; j < n; j++ {
			if !v[j] && pieces[j][0] == arr[i] {
				v[j] = true
				break
			}
		}
		if j == n {
			return false
		}
		p := pieces[j]
		for _, vp := range p {
			if arr[i] != vp {
				return false
			}
			i++
		}
	}
	return true
}
