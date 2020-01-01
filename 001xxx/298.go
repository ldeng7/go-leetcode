func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	o, l, q := 0, len(status), make([]int, 0, 32)
	m := make([]bool, l)
	for _, i := range initialBoxes {
		if 1 == status[i] {
			q = append(q, i)
		} else {
			m[i] = true
		}
	}
	for 0 != len(q) {
		i0 := q[0]
		o += candies[i0]
		ar := keys[i0]
		for _, i := range ar {
			if 0 == status[i] && m[i] {
				q = append(q, i)
			}
			status[i] = 1
		}
		ar = containedBoxes[i0]
		for _, i := range ar {
			if 1 == status[i] {
				q = append(q, i)
			} else {
				m[i] = true
			}
		}
		q = q[1:]
	}
	return o
}
