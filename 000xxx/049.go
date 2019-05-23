import "sort"

func sortedString(s string) string {
	ints := make([]int, len(s))
	for i, c := range s {
		ints[i] = int(c)
	}
	sort.Ints(ints)
	rs := make([]rune, len(s))
	for i, c := range ints {
		rs[i] = rune(c)
	}
	return string(rs)
}

func groupAnagrams(strs []string) [][]string {
	out := [][]string{}
	m := map[string]uint{}
	for _, s := range strs {
		h := sortedString(s)
		if i, ok := m[h]; ok {
			out[i] = append(out[i], s)
		} else {
			m[h] = uint(len(out))
			out = append(out, []string{s})
		}
	}
	return out
}
