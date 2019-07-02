func removeStones(stones [][]int) int {
	l, c := len(stones), 0
	t := make([]bool, l)
	var cal func(int, int, int)
	cal = func(i, y0, x0 int) {
		t[i] = true
		for j := 0; j < l; j++ {
			y, x := stones[j][0], stones[j][1]
			if (y0 == y || x0 == x) && !t[j] {
				cal(j, y, x)
			}
		}
	}

	for i := 0; i < l; i++ {
		if !t[i] {
			cal(i, stones[i][0], stones[i][1])
			c++
		}
	}
	return l - c
}
