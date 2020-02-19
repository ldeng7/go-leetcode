func numberOfSteps(num int) int {
	o := 0
	for ; 0 != num; num >>= 1 {
		if num&1 == 1 {
			o += 2
		} else {
			o++
		}
	}
	return o - 1
}
