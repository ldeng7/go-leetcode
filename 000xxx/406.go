import "sort"

type People [][]int

func (p People) Len() int {
	return len(p)
}

func (p People) Less(i, j int) bool {
	return p[i][0] > p[j][0] || (p[i][0] == p[j][0] && p[i][1] < p[j][1])
}

func (p People) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func reconstructQueue(people [][]int) [][]int {
	sort.Sort(People(people))
	out := [][]int{}
	for _, p := range people {
		out = append(out, p)
		if p[1] != len(out)-1 {
			copy(out[p[1]+1:], out[p[1]:len(out)-1])
			out[p[1]] = p
		}
	}
	return out
}
