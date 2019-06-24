import "math"

var dirs = [8][2]int{{-1, -2}, {-2, -1}, {-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}}

func knightProbability(N int, K int, r int, c int) float64 {
	if K == 0 {
		return 1
	}
	t := make([][]float64, N)
	for i := 0; i < N; i++ {
		t[i] = make([]float64, N)
		for j := 0; j < N; j++ {
			t[i][j] = 1
		}
	}
	for k := 0; k < K; k++ {
		t1 := make([][]float64, N)
		for i := 0; i < N; i++ {
			t1[i] = make([]float64, N)
		}
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				for _, d := range dirs {
					y, x := i+d[0], j+d[1]
					if y < 0 || y >= N || x < 0 || x >= N {
						continue
					}
					t1[i][j] += t[y][x]
				}
			}
		}
		t = t1
	}
	return t[r][c] / math.Pow(8, float64(K))
}
