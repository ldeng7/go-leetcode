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

func luckyNumbers(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	o := make([]int, 0, max(m, n))
	set := make(map[int]bool, m)
	for _, row := range matrix {
		a := row[0]
		for j := 1; j < n; j++ {
			a = min(a, row[j])
		}
		set[a] = true
	}
	for i := 0; i < n; i++ {
		a := matrix[0][i]
		for j := 1; j < m; j++ {
			a = max(a, matrix[j][i])
		}
		if set[a] {
			o = append(o, a)
		}
	}
	return o
}
