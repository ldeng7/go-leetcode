func minDeletionSize(A []string) int {
	cnt := 0
	for j := 0; j < len(A[0]); j++ {
		for i := 1; i < len(A); i++ {
			if A[i-1][j] > A[i][j] {
				cnt++
				break
			}
		}
	}
	return cnt
}
