import "sort"

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || (a[0] == b[0] && a[1] > b[1])
	})
	t := []int{}
	for i := 0; i < len(envelopes); i++ {
		h := envelopes[i][1]
		k := sort.Search(len(t), func(j int) bool {
			return t[j] >= h
		})
		if k == len(t) {
			t = append(t, h)
		} else {
			t[k] = h
		}
	}
	return len(t)
}
