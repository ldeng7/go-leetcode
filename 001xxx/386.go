func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	o := n << 1
	m := map[int]int{}
	for _, s := range reservedSeats {
		m[s[0]] |= 1 << (s[1] - 1)
	}
	for _, v := range m {
		if v&0x01fe != 0 {
			o--
		}
		if (v&0x0078 != 0) && (v&0x001e != 0) && (v&0x01e0 != 0) {
			o--
		}
	}
	return o
}
