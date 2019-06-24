func brokenCalc(X int, Y int) int {
	o := 0
	for X < Y {
		if Y&1 == 0 {
			Y, o = Y>>1, o+1
		} else {
			Y, o = (Y+1)>>1, o+2
		}
	}
	return o + X - Y
}
