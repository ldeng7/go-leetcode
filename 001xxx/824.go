func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minSideJumps(obstacles []int) int {
	ar := [4]int{0, 1, 0, 1}
	for i := 1; i < len(obstacles); i++ {
		t := obstacles[i]
		ar[t] = 1000000
		if t != 1 {
			ar[1] = min(ar[1], min(ar[2], ar[3])+1)
		}
		if t != 2 {
			ar[2] = min(ar[2], min(ar[3], ar[1])+1)
		}
		if t != 3 {
			ar[3] = min(ar[3], min(ar[1], ar[2])+1)
		}
	}
	return min(ar[1], min(ar[2], ar[3]))
}
