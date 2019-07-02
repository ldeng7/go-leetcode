import "sort"

func highFive(items [][]int) [][]int {
	m := [1001][]int{}
	for _, item := range items {
		id, s := item[0], item[1]
		m[id] = append(m[id], s)
	}
	o := [][]int{}
	for id, ss := range m {
		if nil == ss {
			continue
		}
		sort.Ints(ss)
		e := len(ss) - 1
		o = append(o, []int{id, (ss[e] + ss[e-1] + ss[e-2] + ss[e-3] + ss[e-4]) / 5})
	}
	return o
}
