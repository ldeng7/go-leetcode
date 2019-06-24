import "sort"

func frequencySort(s string) string {
	m, bs := map[byte]int{}, []byte(s)
	for _, b := range bs {
		m[b]++
	}
	sort.Slice(bs, func(i, j int) bool {
		a, b := bs[i], bs[j]
		return m[a] > m[b] || (m[a] == m[b] && a < b)
	})
	return string(bs)
}
