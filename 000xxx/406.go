import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})
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

