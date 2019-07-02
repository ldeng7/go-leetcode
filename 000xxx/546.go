func removeBoxes(boxes []int) int {
	t := [100][100][100]int{}
	var cal func(i, j, k int) int
	cal = func(i, j, k int) int {
		if j < i {
			return 0
		} else if t[i][j][k] > 0 {
			return t[i][j][k]
		}
		o := (k+1)*(k+1) + cal(i+1, j, 0)
		for h := i + 1; h <= j; h++ {
			if boxes[h] == boxes[i] {
				if o1 := cal(i+1, h-1, 0) + cal(h, j, k+1); o1 > o {
					o = o1
				}
			}
		}
		t[i][j][k] = o
		return o
	}
	return cal(0, len(boxes)-1, 0)
}
