import "sort"

func merge(intervals [][]int) [][]int {
	out := [][]int{}
	l := len(intervals)
	bs, es := make([]int, l), make([]int, l)
	for i, interval := range intervals {
		bs[i], es[i] = interval[0], interval[1]
	}
	sort.Ints(bs)
	sort.Ints(es)
	for i, j := 0, 0; i < l; i++ {
		if i == l-1 || bs[i+1] > es[i] {
			out = append(out, []int{bs[j], es[i]})
			j = i + 1
		}
	}
	return out
}
