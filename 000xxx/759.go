import "sort"

func employeeFreeTime(schedule [][][]int) [][]int {
	m := map[int]int{}
	c := 0
	out := [][]int{}
	for _, e := range schedule {
		for _, itvl := range e {
			m[itvl[0]]++
			m[itvl[1]]--
		}
	}
	ar := []int{}
	for i, _ := range m {
		ar = append(ar, i)
	}
	sort.Ints(ar)
	for _, i := range ar {
		c += m[i]
		if c == 0 {
			out = append(out, []int{i, 0})
		} else if 0 != len(out) && 0 == out[len(out)-1][1] {
			out[len(out)-1][1] = i
		}
	}
	if 0 != len(out) {
		out = out[:len(out)-1]
	}
	return out
}
