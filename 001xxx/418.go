import (
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	fm := map[string]bool{}
	m := map[int]map[string]int{}
	for _, order := range orders {
		i, _ := strconv.Atoi(order[1])
		f := order[2]
		fm[f] = true
		if mi, ok := m[i]; !ok {
			m[i] = map[string]int{f: 1}
		} else {
			mi[f]++
		}
	}

	l := len(fm) + 1
	h := make([]string, 1, l)
	h[0] = "Table"
	for f, _ := range fm {
		h = append(h, f)
	}
	sort.Strings(h[1:])
	o := make([][]string, 1, len(m)+1)
	o[0] = h

	indices := make([]int, 0, len(m))
	for i, _ := range m {
		indices = append(indices, i)
	}
	sort.Ints(indices)
	for _, index := range indices {
		strs := make([]string, l)
		strs[0] = strconv.Itoa(index)
		mi := m[index]
		for i := 1; i < l; i++ {
			strs[i] = strconv.Itoa(mi[h[i]])
		}
		o = append(o, strs)
	}
	return o
}
