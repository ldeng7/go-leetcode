func rotateTheBox(box [][]byte) [][]byte {
	m, n := len(box), len(box[0])
	o := make([][]byte, n)
	for i := 0; i < n; i++ {
		ar := make([]byte, m)
		for j := 0; j < m; j++ {
			ar[j] = '.'
		}
		o[i] = ar
	}
	for i := 0; i < m; i++ {
		k := n - 1
		for j := n - 1; j >= 0; j-- {
			switch box[i][j] {
			case '*':
				o[j][m-i-1] = '*'
				k = j - 1
			case '#':
				o[k][m-i-1] = '#'
				k--
			}
		}
	}
	return o
}
