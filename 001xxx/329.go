import "sort"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func diagonalSort(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	l := min(m, n)
	for i := n - 1; i >= 0; i-- {
		ar := make([]int, 0, l)
		r, c := 0, i
		for r < m && c < n {
			ar = append(ar, mat[r][c])
			r, c = r+1, c+1
		}
		sort.Ints(ar)
		r, c, k := 0, i, 0
		for r < m && c < n {
			mat[r][c] = ar[k]
			r, c, k = r+1, c+1, k+1
		}
	}
	for i := 1; i < m; i++ {
		ar := make([]int, 0, l)
		r, c := i, 0
		for r < m && c < n {
			ar = append(ar, mat[r][c])
			r, c = r+1, c+1
		}
		sort.Ints(ar)
		r, c, k := i, 0, 0
		for r < m && c < n {
			mat[r][c] = ar[k]
			r, c, k = r+1, c+1, k+1
		}
	}
	return mat
}
