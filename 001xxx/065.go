import (
	"sort"
	"strings"
)

func indexPairs(text string, words []string) [][]int {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	ie := len(text) - len(words[0]) + 1
	out := [][]int{}
	for i := 0; i < ie; i++ {
		s := text[i:]
		for _, w := range words {
			if strings.HasPrefix(s, w) {
				out = append(out, []int{i, i + len(w) - 1})
			}
		}
	}
	return out
}
