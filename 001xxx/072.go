func maxEqualRowsAfterFlips(matrix [][]int) int {
	m := map[string]int{}
	l, n := len(matrix), len(matrix[0])
	for i := 0; i < l; i++ {
		bs1, bs2 := make([]byte, n), make([]byte, n)
		for j := 0; j < n; j++ {
			bs1[j] = '0' + byte(matrix[i][j])
			bs2[j] = '0' + byte(matrix[i][j]^1)
		}
		m[string(bs1)]++
		m[string(bs2)]++
	}
	o := 0
	for _, c := range m {
		if c > o {
			o = c
		}
	}
	return o
}
