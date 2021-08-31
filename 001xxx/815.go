func maxHappyGroups(batchSize int, groups []int) int {
	cs := make([]int, batchSize)
	o := 0
	for _, g := range groups {
		g %= batchSize
		if g == 0 {
			o++
		} else if cs[batchSize-g] > 0 {
			cs[batchSize-g]--
			o++
		} else {
			cs[g]++
		}
	}

	m := map[int]int{}
	key := func() int {
		k := 0
		for i := 0; i < batchSize; i++ {
			k = (k << 8) | cs[i]
		}
		return k
	}

	var cal func(int) int
	cal = func(l int) int {
		if o, ok := m[key()]; ok {
			return o
		}
		o := 0
		for i := 1; i < batchSize; i++ {
			cs[i]--
			if cs[i] >= 0 {
				t := cal((batchSize + l - i) % batchSize)
				if l == 0 {
					t++
				}
				if t > o {
					o = t
				}
			}
			cs[i]++
		}
		m[key()] = o
		return o
	}

	return o + cal(0)
}
