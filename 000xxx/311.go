func multiply(A [][]int, B [][]int) [][]int {
	ha, wa, wb := len(A), len(A[0]), len(B[0])
	out := make([][]int, ha)
	for i := 0; i < ha; i++ {
		out[i] = make([]int, wb)
	}
	for i := 0; i < ha; i++ {
		for j := 0; j < wa; j++ {
			a := A[i][j]
			if a == 0 {
				continue
			}
			br := B[j]
			for k := 0; k < wb; k++ {
				if br[k] == 0 {
					continue
				}
				out[i][k] += a * br[k]
			}
		}
	}
	return out
}
