func minSwaps(grid [][]int) int {
	l := len(grid)
	t := make([]int, l)
	for i := 0; i < l; i++ {
		t[i] = -1
	}
	for i := 0; i < l; i++ {
		k := 0
		for j := l - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				k = j
				break
			}
		}
		for ; k < l && t[k] != -1; k++ {
		}
		if k == l {
			return -1
		} else {
			t[k] = i
		}
	}
	o := 0
	for j := 0; j < l; j++ {
		for i := 0; i < l-1; i++ {
			if t[i] > t[i+1] {
				t[i], t[i+1], o = t[i+1], t[i], o+1
			}
		}
	}
	return o
}
