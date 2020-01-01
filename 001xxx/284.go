import "math"

var dirs = [5][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {0, 0}}

func minFlips(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	trans := func(i, j int) {
		for _, d := range dirs {
			i1, j1 := i+d[0], j+d[1]
			if i1 >= 0 && i1 < m && j1 >= 0 && j1 < n {
				mat[i1][j1] ^= 1
			}
		}
	}

	o := math.MaxInt64
	var cal func(int, int)
	cal = func(p, c int) {
		if p == m*n {
			row, i := mat[m-1], 0
			for ; i < n; i++ {
				if row[i] != 0 {
					break
				}
			}
			if i == n && c < o {
				o = c
			}
		} else {
			i, j := p/n, p%n
			if i == 0 {
				cal(p+1, c)
				trans(i, j)
				cal(p+1, c+1)
				trans(i, j)
			} else if mat[i-1][j] == 0 {
				cal(p+1, c)
			} else {
				trans(i, j)
				cal(p+1, c+1)
				trans(i, j)
			}
		}
	}

	cal(0, 0)
	if o != math.MaxInt64 {
		return o
	}
	return -1
}
