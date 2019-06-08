import "bytes"

func maskPII(S string) string {
	buf := bytes.Buffer{}
	bs := []byte{}
	for _, c := range []byte(S) {
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			bs = append(bs, c)
		} else if c >= 'A' && c <= 'Z' {
			bs = append(bs, c+32)
		} else if c == '@' {
			buf.WriteByte(bs[0])
			buf.WriteString("*****")
			buf.WriteByte(bs[len(bs)-1])
			buf.WriteByte('@')
			bs = []byte{}
		} else if c == '.' {
			buf.Write(bs)
			buf.WriteByte('.')
			bs = []byte{}
		}
	}
	if buf.Len() != 0 {
		buf.Write(bs)
		return buf.String()
	}
	l := len(bs)
	if l > 10 {
		buf.WriteByte('+')
		for i := 0; i < l-10; i++ {
			buf.WriteByte('*')
		}
		buf.WriteByte('-')
	}
	buf.WriteString("***-***-")
	buf.Write(bs[l-4:])
	return buf.String()
}
