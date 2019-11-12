func oddCells(n int, m int, indices [][]int) int {
	rows, cols := make([]int, n), make([]int, m)
	for _, i := range indices {
		rows[i[0]]++
		cols[i[1]]++
	}
	cr, cc := 0, 0
	for _, i := range rows {
		if i&1 == 1 {
			cr++
		}
	}
	for _, i := range cols {
		if i&1 == 1 {
			cc++
		}
	}
	return cr*(m-cc) + cc*(n-cr)
}
