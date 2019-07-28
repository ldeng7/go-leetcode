import "strings"

func reverseWords(s string) string {
	es := strings.Split(s, " ")
	es1 := make([]string, 0, len(es))
	for i := len(es) - 1; i >= 0; i-- {
		if 0 != len(es[i]) {
			es1 = append(es1, es[i])
		}
	}
	return strings.Join(es1, " ")
}
