import (
	"sort"
	"strings"
)

type AlphaLogs struct {
	logs []string
	keys [][2]string
}

func (al *AlphaLogs) Len() int { return len(al.logs) }
func (al *AlphaLogs) Less(i, j int) bool {
	t1, t2, c1, c2 := al.logs[i][0], al.logs[j][0], al.keys[i][1], al.keys[j][1]
	return c1 < c2 || (c1 == c2 && t1 < t2)
}
func (al *AlphaLogs) Swap(i, j int) {
	al.logs[i], al.logs[j] = al.logs[j], al.logs[i]
	al.keys[i], al.keys[j] = al.keys[j], al.keys[i]
}

func reorderLogFiles(logs []string) []string {
	l := len(logs)
	ar1, ar2 := make([]string, 0, l), make([]string, 0, l)
	for _, s := range logs {
		if c := s[len(s)-1]; c >= 'a' && c <= 'z' {
			ar1 = append(ar1, s)
		} else {
			ar2 = append(ar2, s)
		}
	}

	ks := make([][2]string, len(ar1))
	for i, s := range ar1 {
		j := strings.IndexByte(s, ' ')
		ks[i] = [2]string{s[:j], s[j+1:]}
	}
	al := &AlphaLogs{ar1, ks}
	sort.Sort(al)

	out := make([]string, l)
	copy(out, ar1)
	copy(out[len(ar1):], ar2)
	return out
}
