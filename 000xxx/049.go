import "sort"

func groupAnagrams(strs []string) [][]string {
	out := [][]string{}
	m := map[string]uint{}
	for _, s := range strs {
		bs := []byte(s)
		sort.Slice(bs, func(i, j int) bool {
			return bs[i] < bs[j]
		})
		h := string(bs)
		if i, ok := m[h]; ok {
			out[i] = append(out[i], s)
		} else {
			m[h] = uint(len(out))
			out = append(out, []string{s})
		}
	}
	return out
}
