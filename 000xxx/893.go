import "sort"

func numSpecialEquivGroups(A []string) int {
	m := map[string]map[string]bool{}
	for _, s := range A {
		cs0, cs1 := []int{}, []int{}
		for i, c := range s {
			if i&1 == 0 {
				cs0 = append(cs0, int(c))
			} else {
				cs1 = append(cs1, int(c))
			}
		}
		sort.Ints(cs0)
		sort.Ints(cs1)

		bs0, bs1 := make([]byte, len(cs0)), make([]byte, len(cs1))
		for i, c := range cs0 {
			bs0[i] = byte(c)
		}
		for i, c := range cs1 {
			bs1[i] = byte(c)
		}
		s0, s1 := string(bs0), string(bs1)

		if _, ok := m[s0]; ok {
			if _, ok := m[s0][s1]; !ok {
				m[s0][s1] = true
			}
		} else {
			m[s0] = map[string]bool{s1: true}
		}
	}
	cnt := 0
	for _, sm := range m {
		cnt += len(sm)
	}
	return cnt
}

