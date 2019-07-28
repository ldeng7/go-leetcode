func smallestSufficientTeam(reqSkills []string, people [][]string) []int {
	o, t, l := []int{}, []int{}, len(people)
	m, ss := map[string]uint{}, make([]int, l)
	for i, s := range reqSkills {
		m[s] = uint(i)
	}
	for i, p := range people {
		for _, s := range p {
			ss[i] |= 1 << m[s]
		}
	}

	var cal func(int)
	cal = func(i int) {
		if i == (1<<uint(len(reqSkills)))-1 {
			if 0 == len(o) || len(o) > len(t) {
				o = make([]int, len(t))
				copy(o, t)
			}
			return
		} else if 0 != len(o) && len(o) <= len(t) {
			return
		}
		var k uint = 0
		for ((i >> k) & 1) == 1 {
			k++
		}
		for j, s := range ss {
			if ((s >> k) & 1) == 1 {
				t = append(t, j)
				cal(i | s)
				t = t[:len(t)-1]
			}
		}
	}
	cal(0)
	return o
}
