import (
	"bytes"
	"strings"
)

var vowels map[byte]bool = map[byte]bool{
	'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
}

func toGoatLatin(S string) string {
	words := strings.Split(S, " ")
	buf := &bytes.Buffer{}
	c := 2
	for _, w := range words {
		buf.WriteByte(' ')
		if vowels[w[0]] {
			buf.WriteString(w)
		} else {
			buf.WriteString(w[1:])
			buf.WriteByte(w[0])
		}
		buf.WriteByte('m')
		for i := 0; i < c; i++ {
			buf.WriteByte('a')
		}
		c++
	}
	buf.ReadByte()
	return buf.String()
}
