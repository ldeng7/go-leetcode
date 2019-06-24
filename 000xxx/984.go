import "bytes"

func strWithout3a3b(A int, B int) string {
	buf := &bytes.Buffer{}
	for A > 0 || B > 0 {
		if A > B {
			if A > 1 {
				buf.WriteString("aa")
				A -= 2
			} else {
				buf.WriteByte('a')
				A--
			}
			if B > 0 {
				buf.WriteByte('b')
				B--
			}
		} else if A < B {
			if B > 1 {
				buf.WriteString("bb")
				B -= 2
			} else {
				buf.WriteByte('b')
				B--
			}
			if A > 0 {
				buf.WriteByte('a')
				A--
			}
		} else {
			buf.WriteString("ab")
			A--
			B--
		}
	}
	return buf.String()
}
