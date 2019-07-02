func prisonAfterNDays(cells []int, N int) []int {
	N %= 14
	if N == 0 {
		N = 14
	}
	for i := 0; i < N; i++ {
		t := make([]int, 8)
		for j := 1; j < 7; j++ {
			if cells[j-1] == cells[j+1] {
				t[j] = 1
			}
		}
		cells = t
	}
	return cells
}
