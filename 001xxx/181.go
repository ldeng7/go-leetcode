import (
	"sort"
	"strings"
)

func beforeAndAfterPuzzles(phrases []string) []string {
	l := len(phrases)
	words := make([][]string, l)
	head := map[string][]int{}
	set := map[string]bool{}
	for i, p := range phrases {
		ar := strings.Split(p, " ")
		s := ar[0]
		words[i] = ar
		head[s] = append(head[s], i)
	}
	for i := 0; i < l; i++ {
		ar := words[i]
		es0 := strings.Join(ar, " ")
		s := ar[len(ar)-1]
		for _, k := range head[s] {
			if k == i {
				continue
			}
			es := make([]string, 0, 2)
			if 0 != len(es0) {
				es = append(es, es0)
			}
			es1 := strings.Join(words[k][1:], " ")
			if 0 != len(es1) {
				es = append(es, es1)
			}
			set[strings.Join(es, " ")] = true
		}
	}
	o := make([]string, 0, len(set))
	for s, _ := range set {
		o = append(o, s)
	}
	sort.Strings(o)
	return o
}
