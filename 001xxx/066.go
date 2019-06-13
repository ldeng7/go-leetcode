func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func assignBikes(workers [][]int, bikes [][]int) int {
	lw, lb := uint8(len(workers)), uint8(len(bikes))
	var ibse uint16 = (1 << lb) - 1
	t := make([]int, ibse)
	var cal func(uint8, uint16) int
	cal = func(iw uint8, ibs uint16) int {
		if iw == lw {
			return 0
		} else if t[ibs] != 0 {
			return t[ibs]
		}

		t1 := 20000
		for i := uint8(0); i < lb; i++ {
			if (ibs & (1 << i)) == 0 {
				w, b := workers[iw], bikes[i]
				t2 := cal(iw+1, ibs|(1<<i)) + abs(w[0]-b[0]) + abs(w[1]-b[1])
				if t2 < t1 {
					t1 = t2
				}
			}
		}
		t[ibs] = t1
		return t1
	}
	return cal(0, 0)
}
