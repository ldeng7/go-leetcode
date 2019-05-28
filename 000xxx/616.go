import "bytes"

func addBoldTag(s string, dict []string) string {
	if 0 == len(s) {
		return ""
	}
	bold := make([]bool, len(s))
	for _, w := range dict {
		for i := 0; i <= len(s)-len(w); i++ {
			j := 0
			for ; j < len(w); j++ {
				if s[i+j] != w[j] {
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
	buf.WriteByte(s[0])
	for i := 1; i < len(s); i++ {
		if bold[i] && !bold[i-1] {
			buf.WriteString("<b>")
		} else if !bold[i] && bold[i-1] {
			buf.WriteString("</b>")
		}
		buf.WriteByte(s[i])
	}
	if bold[len(s)-1] {
		buf.WriteString("</b>")
	}
	return buf.String()
}
