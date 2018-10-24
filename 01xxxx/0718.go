func findLength(a []int, b []int) int {
	out := 0
	t := make([][]int, len(a)+1)
	for i := 0; i <= len(a); i++ {
		t[i] = make([]int, len(b)+1)
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				t[i+1][j+1] = t[i][j] + 1
			} else {
				t[i+1][j+1] = 0
			}
			if t[i+1][j+1] > out {
				out = t[i+1][j+1]
			}
		}
	}
	return out
}
