import "sort"

func isNStraightHand(hand []int, W int) bool {
	m := map[int]int{}
	ks := []int{}
	for _, h := range hand {
		if _, ok := m[h]; !ok {
			ks = append(ks, h)
		}
		m[h]++
	}
	sort.Ints(ks)
	for _, h := range ks {
		c := m[h]
		if c == 0 {
			continue
		}
		for j := h; j < h+W; j++ {
			c1, _ := m[j]
			if c1 < c {
				return false
			}
			m[j] = c1 - c
		}
	}
	return true
}
