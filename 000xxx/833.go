import (
	"sort"
)

func findReplaceString(S string, indexes []int, sources []string, targets []string) string {
	n := len(S)
	ps := [][2]int{}
	for i, idx := range indexes {
		r := idx + len(sources[i])
		if r <= n && S[idx:r] == sources[i] {
			ps = append(ps, [2]int{idx, i})
		}
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i][0] > ps[j][0]
	})
	for _, p := range ps {
		l, i := p[0], p[1]
		S = S[:l] + targets[i] + S[l+len(sources[i]):]
	}
	return S
}
