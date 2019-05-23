import "strings"

func uncommonFromSentences(A string, B string) []string {
	out := []string{}
	ma, mb := map[string]int{}, map[string]int{}
	for _, s := range strings.Split(A, " ") {
		if _, ok := ma[s]; ok {
			ma[s]++
		} else if 0 != len(s) {
			ma[s] = 1
		}
	}
	for _, s := range strings.Split(B, " ") {
		if _, ok := mb[s]; ok {
			mb[s]++
		} else if 0 != len(s) {
			mb[s] = 1
		}
	}
	for s, c := range ma {
		if _, ok := mb[s]; !ok && 1 == c {
			out = append(out, s)
		}
	}
	for s, c := range mb {
		if _, ok := ma[s]; !ok && 1 == c {
			out = append(out, s)
		}
	}
	return out
}
