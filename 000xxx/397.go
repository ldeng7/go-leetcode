func integerReplacement(n int) int {
	o := 0
	for n > 1 {
		o++
		if n&1 != 0 {
			if (n&2 != 0) && (n != 3) {
				n++
			} else {
				n--
			}
		} else {
			n >>= 1
		}
	}
	return o
}
