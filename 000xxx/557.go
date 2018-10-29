import "strings"

func reverseWords(s string) string {
	es := strings.Split(s, " ")
	for i, e := range es {
		bs := make([]byte, len(e))
		for j := len(e) - 1; j >= 0; j-- {
			bs[len(e)-j-1] = e[j]
		}
		es[i] = string(bs)
	}
	return strings.Join(es, " ")
}
