func minScoreTriangulation(A []int) int {
	l := len(A)
	t := make([][]int, l)
	t[0] = make([]int, l)
	for i := 1; i < l; i++ {
		t[i] = make([]int, l)
		t[i-1][i] = 0
		for j := i - 2; j >= 0; j-- {
			for k := j + 1; k < i; k++ {
				t1 := t[j][k] + t[k][i] + A[i]*A[j]*A[k]
				if t[j][i] == 0 || t[j][i] > t1 {
					t[j][i] = t1
				}
			}
		}
	}
	return t[0][l-1]
}
