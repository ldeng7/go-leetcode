import "sort"

func peopleIndexes(favoriteCompanies [][]string) []int {
	l := len(favoriteCompanies)
	m := map[string]int{}
	hs := make([][]int, l)
	for i, ss := range favoriteCompanies {
		h := make([]int, len(ss))
		for j, s := range ss {
			k, ok := m[s]
			if !ok {
				k = len(m)
				m[s] = k
			}
			h[j] = k
		}
		sort.Ints(h)
		hs[i] = h
	}

	isSubset := func(i, j int) bool {
		if i == j {
			return false
		}
		ha, hb := hs[i], hs[j]
		la, lb := len(ha), len(hb)
		ia, ib := 0, 0
		for ia < la && ib < lb {
			ka, kb := ha[ia], hb[ib]
			if ka == kb {
				ia++
				ib++
			} else if ka > kb {
				ib++
			} else {
				return false
			}
		}
		return ia == la
	}

	o := make([]int, 0, 64)
	for i := 0; i < l; i++ {
		j := 0
		for ; j < l && !isSubset(i, j); j++ {
		}
		if j == l {
			o = append(o, i)
		}
	}
	return o
}
