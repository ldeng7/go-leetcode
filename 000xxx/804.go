import "bytes"

var table [26]string = [26]string{
	".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--",
	"-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
}

func uniqueMorseRepresentations(words []string) int {
	m := map[string]bool{}
	for _, w := range words {
		buf := &bytes.Buffer{}
		for i := 0; i < len(w); i++ {
			buf.WriteString(table[w[i]-'a'])
		}
		m[buf.String()] = true
	}
	return len(m)
}
