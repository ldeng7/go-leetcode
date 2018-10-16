func generate(numRows int) [][]int {
	if 0 == numRows {
		return [][]int{}
	}
	out := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		up := out[i-1]
		row := make([]int, i+1)
		row[0], row[len(row)-1] = 1, 1
		for j := 1; j < len(row)-1; j++ {
			row[j] = up[j-1] + up[j]
		}
		out = append(out, row)
	}
	return out
}
