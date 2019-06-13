import "strings"

func findOcurrences(text string, first string, second string) []string {
	k := first + " " + second
	out := []string{}
	for {
		i := strings.Index(text, k)
		j := i + len(k) + 1
		if -1 == i || j >= len(text) {
			break
		}
		text = text[j:]
		i = strings.IndexByte(text, ' ')
		if i >= 0 {
			out = append(out, text[:i])
		} else {
			out = append(out, text)
		}
	}
	return out
}
