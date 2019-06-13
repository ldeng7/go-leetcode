func predictPartyVictory(senate string) string {
	l := len(senate)
	q1, q2 := []int{}, []int{}
	for i := 0; i < l; i++ {
		if senate[i] == 'R' {
			q1 = append(q1, i)
		} else {
			q2 = append(q2, i)
		}
	}
	for 0 != len(q1) && 0 != len(q2) {
		i, j := q1[0], q2[0]
		q1, q2 = q1[1:], q2[1:]
		if i < j {
			q1 = append(q1, i+l)
		} else {
			q2 = append(q2, j+l)
		}
	}
	if len(q1) > len(q2) {
		return "Radiant"
	}
	return "Dire"
}
