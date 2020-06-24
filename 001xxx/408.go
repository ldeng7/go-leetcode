import (
	"sort"
	"strings"
)

func stringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	l := len(words)
	o := make([]string, 0, l)
	for i := 1; i < l; i++ {
		s := words[i]
		for j := 0; j < i; j++ {
			if sp := words[j]; strings.Contains(sp, s) {
				o = append(o, s)
				break
			}
		}
	}
	return o
}
