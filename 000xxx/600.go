func findIntegers(num int) int {
	o, p, k := 0, 0, 31
	f := [32]int{}
	f[0], f[1] = 1, 2
	for i := 2; i < 31; i++ {
		f[i] = f[i-2] + f[i-1]
	}
	for k >= 0 {
		if num&(1<<uint(k)) != 0 {
			o += f[k]
			if 0 != p {
				return o
			}
			p = 1
		} else {
			p = 0
		}
		k--
	}
	return o + 1
}
