import "math"

func maxSumSubmatrix(matrix [][]int, k int) int {
	if 0 == len(matrix) || 0 == len(matrix[0]) {
		return 0
	}
	o, m, n := math.MinInt64, len(matrix), len(matrix[0])
	ss := make([][]int, m)
	for i := 0; i < m; i++ {
		ss[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			t := matrix[i][j]
			if i > 0 {
				t += ss[i-1][j]
			}
			if j > 0 {
				t += ss[i][j-1]
			}
			if i > 0 && j > 0 {
				t -= ss[i-1][j-1]
			}
			ss[i][j] = t
			for r := 0; r <= i; r++ {
				for c := 0; c <= j; c++ {
					s := ss[i][j]
					if r > 0 {
						s -= ss[r-1][j]
					}
					if c > 0 {
						s -= ss[i][c-1]
					}
					if r > 0 && c > 0 {
						s += ss[r-1][c-1]
					}
					if s <= k && s > o {
						o = s
					}
				}
			}
		}
	}
	return o
}
