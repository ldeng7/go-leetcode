func numSubmatrixSumTarget(matrix [][]int, target int) int {
	m, n := len(matrix), len(matrix[0])
	o, t := 0, map[int]int{}
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}
	for k := 0; k < n; k++ {
		for j := k; j < n; j++ {
			t = map[int]int{0: 1}
			s := 0
			for i := 0; i < m; i++ {
				if k > 0 {
					s += matrix[i][j] - matrix[i][k-1]
				} else {
					s += matrix[i][j]
				}
				if v, ok := t[s-target]; ok {
					o += v
				}
				t[s]++
			}
		}
	}
	return o
}
