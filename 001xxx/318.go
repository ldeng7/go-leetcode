func minFlips(a int, b int, c int) int {
	o := 0
	for c != 0 || a != 0 || b != 0 {
		a1, b1, c1 := a&1, b&1, c&1
		a, b, c = a>>1, b>>1, c>>1
		if (a1 | b1) == c1 {
			continue
		}
		if c1 == 1 {
			o++
		} else {
			if a1 == 1 {
				o++
			}
			if b1 == 1 {
				o++
			}
		}
	}
	return o
}
