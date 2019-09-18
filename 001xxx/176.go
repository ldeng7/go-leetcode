func findNumOfValidWords(words []string, puzzles []string) []int {
	o := make([]int, len(puzzles))
	m := map[uint32]int{}
	var bits uint32
	var cal func(int, string) int
	cal = func(i int, s string) int {
		if i == len(s) {
			return m[bits]
		}
		t := bits
		c := cal(i+1, s)
		bits |= 1 << uint8(s[i]-'a')
		c += cal(i+1, s)
		bits = t
		return c
	}

	for _, w := range words {
		bits = 0
		for i := 0; i < len(w); i++ {
			bits |= 1 << uint8(w[i]-'a')
		}
		m[bits]++
	}
	for i, s := range puzzles {
		bits = 0
		bits |= 1 << uint8(s[0]-'a')
		o[i] = cal(1, s)
	}
	return o
}
