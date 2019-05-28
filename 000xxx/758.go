import "bytes"

func boldWords(words []string, S string) string {
	if 0 == len(S) {
		return ""
	}
	bold := make([]bool, len(S))
	for _, w := range words {
		for i := 0; i <= len(S)-len(w); i++ {
			j := 0
			for ; j < len(w); j++ {
				if S[i+j] != w[j] {
					break
				}
			}
			if j == len(w) {
				for j := 0; j < len(w); j++ {
					bold[i+j] = true
				}
			}
		}
	}

	buf := &bytes.Buffer{}
	if bold[0] {
		buf.WriteString("<b>")
	}
	buf.WriteByte(S[0])
	for i := 1; i < len(S); i++ {
		if bold[i] && !bold[i-1] {
			buf.WriteString("<b>")
		} else if !bold[i] && bold[i-1] {
			buf.WriteString("</b>")
		}
		buf.WriteByte(S[i])
	}
	if bold[len(S)-1] {
		buf.WriteString("</b>")
	}
	return buf.String()
}
