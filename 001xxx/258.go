import (
	"sort"
	"strings"
)

func generateSentences(synonyms [][]string, text string) []string {
	ar := strings.Split(text, " ")
	set := map[string]bool{}
	m := map[string]map[string]bool{}

	var cal func(int)
	cal = func(ib int) {
		for i := ib; i < len(ar); i++ {
			s := ar[i]
			if nil == m[s] {
				continue
			}
			for s1, _ := range m[s] {
				if s1 == s {
					continue
				}
				ar[i] = s1
				set[strings.Join(ar, " ")] = true
				cal(i + 1)
			}
			ar[i] = s
		}
	}

	for _, syn := range synonyms {
		s, t := syn[0], syn[1]
		ms, mt := m[s], m[t]
		if nil != ms {
			m[t], ms[t] = ms, true
		} else if nil != mt {
			m[s], mt[s] = mt, true
		} else {
			set := map[string]bool{s: true, t: true}
			m[s], m[t] = set, set
		}
	}

	cal(0)
	set[text] = true
	o := make([]string, 0, len(set))
	for s, _ := range set {
		o = append(o, s)
	}
	sort.Strings(o)
	return o
}
