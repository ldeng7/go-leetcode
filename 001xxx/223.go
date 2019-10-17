const M = 1000000007

func dieSimulator(n int, rollMax []int) int {
	t := [5001][6][16]int{}
	ss := [6]int{}
	for i := 0; i < 6; i++ {
		t[0][i][1], ss[i] = 1, 1
	}
	s := 6
	for i := 1; i < n; i++ {
		for j := 0; j < 6; j++ {
			t[i][j][1] = s - ss[j]
			ss[j] = t[i][j][1]
			for k := 2; k <= rollMax[j]; k++ {
				t[i][j][k] = t[i-1][j][k-1]
				ss[j] = (ss[j] + t[i][j][k]) % M
			}
		}
		s = 0
		for j := 0; j < 6; j++ {
			s = (s + ss[j]) % M
		}
	}
	return (s + M) % M
}
