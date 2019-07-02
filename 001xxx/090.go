import "sort"

func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	l := len(values)
	t, m := make([][2]int, l), map[int]int{}
	for i := 0; i < l; i++ {
		t[i][0], t[i][1] = values[i], labels[i]
	}
	sort.Slice(t, func(i, j int) bool { return t[i][0] > t[j][0] })
	o, c := 0, 0
	for i := 0; i < l; i++ {
		label := t[i][1]
		if m[label] >= use_limit {
			continue
		}
		o, c = o+t[i][0], c+1
		if c >= num_wanted {
			return o
		}
		m[label]++
	}
	return o
}
