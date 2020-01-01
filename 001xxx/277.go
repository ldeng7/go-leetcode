func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func countSquares(matrix [][]int) int {
	o := 0
	for i, r := range matrix {
		var rp []int
		if i != 0 {
			rp = matrix[i-1]
		}
		for j, v := range r {
			if v == 0 {
				continue
			}
			if i != 0 && j != 0 {
				v = min(min(r[j-1], rp[j]), rp[j-1]) + 1
				r[j] = v
			}
			o += v
		}
	}
	return o
}
