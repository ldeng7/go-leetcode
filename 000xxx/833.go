import (
	"sort"
)

type Pairs [][2]int

func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Less(i, j int) bool { return p[i][0] > p[j][0] }
func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	n := len(S)
	ps := Pairs{}
	for i, idx := range indexes {
		r := idx + len(sources[i])
		if r <= n && S[idx:r] == sources[i] {
			ps = append(ps, [2]int{idx, i})
		}
	}
	sort.Sort(ps)
	for _, p := range ps {
		l, i := p[0], p[1]
		S = S[:l] + targets[i] + S[l+len(sources[i]):]
	}
	return S
}
