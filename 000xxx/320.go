import "fmt"

func generateAbbreviations(word string) []string {
	if 0 == len(word) {
		return []string{""}
	}
	out := []string{fmt.Sprintf("%d", len(word))}
	for i := 0; i < len(word); i++ {
		for _, abbr := range generateAbbreviations(word[i+1:]) {
			if i > 0 {
				out = append(out, fmt.Sprintf("%d%c%s", i, word[i], abbr))
			} else {
				out = append(out, fmt.Sprintf("%c%s", word[i], abbr))
			}
		}
	}
	return out
}
