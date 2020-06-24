import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	es := strings.Split(sentence, " ")
	for i, e := range es {
		if strings.HasPrefix(e, searchWord) {
			return i + 1
		}
	}
	return -1
}
