func wallsAndGates(rooms [][]int) {
	if 0 == len(rooms) || 0 == len(rooms[0]) {
		return
	}
	m, n := len(rooms), len(rooms[0])
	var cal func(i, j, v int)
	cal = func(i, j, v int) {
		if i < 0 || i >= m || j < 0 || j >= n || rooms[i][j] < v {
			return
		}
		rooms[i][j] = v
		cal(i+1, j, v+1)
		cal(i-1, j, v+1)
		cal(i, j+1, v+1)
		cal(i, j-1, v+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if rooms[i][j] == 0 {
				cal(i, j, 0)
			}
		}
	}
}
