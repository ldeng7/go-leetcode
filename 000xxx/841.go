func canVisitAllRooms(rooms [][]int) bool {
	m := map[int]bool{0: true}
	q := []int{0}
	for 0 != len(q) {
		t := q[0]
		q = q[1:]
		for _, k := range rooms[t] {
			if m[k] {
				continue
			}
			m[k] = true
			if len(m) == len(rooms) {
				return true
			}
			q = append(q, k)
		}
	}
	return len(m) == len(rooms)
}
