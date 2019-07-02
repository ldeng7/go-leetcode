import "bytes"

func fullJustify(words []string, maxWidth int) []string {
	o, l, i := []string{}, len(words), 0
	for i < l {
		j, n := i, 0
		for ; j < l && n+len(words[j])+j-i <= maxWidth; j++ {
			n += len(words[j])
		}
		buf := &bytes.Buffer{}
		ns := maxWidth - n
		for k := i; k < j; k++ {
			buf.WriteString(words[k])
			if ns > 0 {
				m := 0
				if j == l {
					if j-k == 1 {
						m = ns
					} else {
						m = 1
					}
				} else {
					t := j - k - 1
					if t > 0 {
						m = ns / t
						if ns%t != 0 {
							m++
						}
					} else {
						m = ns
					}
				}
				ns -= m
				for ; m > 0; m-- {
					buf.WriteByte(' ')
				}
			}
		}
		o, i = append(o, buf.String()), j
	}
	return o
}
