func spiralOrder(matrix [][]int) []int {
	out := []int{}
	if 0 == len(matrix) || 0 == len(matrix[0]) {
		return out
	}
	m, n := len(matrix), len(matrix[0])
	up, down, left, right := 0, m-1, 0, n-1
	for {
		for i := left; i <= right; i++ {
			out = append(out, matrix[up][i])
		}
		up++
		if up > down {
			break
		}
		for i := up; i <= down; i++ {
			out = append(out, matrix[i][right])
		}
		right--
		if left > right {
			break
		}
		for i := right; i >= left; i-- {
			out = append(out, matrix[down][i])
		}
		down--
		if up > down {
			break
		}
		for i := down; i >= up; i-- {
			out = append(out, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}
	return out
}
